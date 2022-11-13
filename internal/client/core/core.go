package core

import (
	"errors"
	"github.com/vleukhin/GophKeeper/internal/client/storage"
	"github.com/vleukhin/GophKeeper/internal/config/client"
	"sync"

	"github.com/fatih/color"
)

type Core struct {
	repo   storage.Repo
	client Client
	cfg    *client.Config
}

var (
	clientUseCase *Core     //nolint:gochecknoglobals // pattern singleton
	once          sync.Once //nolint:gochecknoglobals // pattern singleton
)

func GetApp() *Core {
	once.Do(func() {
		clientUseCase = &Core{}
	})

	return clientUseCase
}

type CoreOptFunc func(*Core)

func SetRepo(r Repo) CoreOptFunc {
	return func(c *Core) {
		c.repo = r
	}
}

func SetAPI(client Client) CoreOptFunc {
	return func(c *Core) {
		c.client = client
	}
}

func SetConfig(cfg *client.Config) CoreOptFunc {
	return func(c *Core) {
		c.cfg = cfg
	}
}

func (uc *Core) InitDB() {
	uc.repo.MigrateDB()
}

var (
	errPasswordCheck = errors.New("invalid password")
	errToken         = errors.New("invalid token")
)

func (c *Core) authorisationCheck(userPassword string) (string, error) {
	if !c.verifyPassword(userPassword) {
		return "", errPasswordCheck
	}
	accessToken, err := c.repo.GetSavedAccessToken()
	if err != nil || accessToken == "" {
		color.Red("User should be logged")

		return "", errToken
	}

	return accessToken, nil
}
