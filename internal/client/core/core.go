package core

import (
	"errors"
	"github.com/vleukhin/GophKeeper/internal/client/api"
	"github.com/vleukhin/GophKeeper/internal/client/storage"
	"github.com/vleukhin/GophKeeper/internal/config/client"
	"github.com/vleukhin/GophKeeper/internal/models"
	"sync"

	"github.com/fatih/color"
)

type GophKeeperClient interface {
	InitDB()

	Register(user *models.User)
	Login(user *models.User)
	Logout()
	Sync(userPassword string)
}

type Core struct {
	storage storage.Repo
	client  api.Client
	cfg     *client.Config
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

type OptFunc func(*Core)

func SetRepo(r storage.Repo) OptFunc {
	return func(c *Core) {
		c.storage = r
	}
}

func SetAPIClient(client api.Client) OptFunc {
	return func(c *Core) {
		c.client = client
	}
}

func SetConfig(cfg *client.Config) OptFunc {
	return func(c *Core) {
		c.cfg = cfg
	}
}

func (c *Core) InitDB() {
	c.storage.MigrateDB()
}

var (
	errPasswordCheck = errors.New("invalid password")
	errToken         = errors.New("invalid token")
)

func (c *Core) authorisationCheck(userPassword string) (string, error) {
	if !c.verifyPassword(userPassword) {
		return "", errPasswordCheck
	}
	accessToken, err := c.storage.GetSavedAccessToken()
	if err != nil || accessToken == "" {
		color.Red("User should be logged")

		return "", errToken
	}

	return accessToken, nil
}
