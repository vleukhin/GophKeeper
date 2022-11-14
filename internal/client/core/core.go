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

type (
	GophKeeperClient interface {
		InitDB()
		Sync(userPassword string)
		SetStorage(r storage.Repo)
		SetAPIClient(client api.Client)
		SetConfig(cfg *client.Config)

		AuthService
		CardsService
		NotesService
	}

	AuthService interface {
		Register(user *models.User)
		Login(user *models.User)
		Logout()
	}
	CardsService interface {
		StoreCard(userPassword string, card *models.Card)
		ShowCard(userPassword, cardID string)
		DelCard(userPassword, cardID string)
	}
	NotesService interface {
		StoreNote(userPassword string, note *models.Note)
		ShowNote(userPassword, noteID string)
		DelNote(userPassword, noteID string)
	}
)

type Core struct {
	storage storage.Repo
	client  api.Client
	cfg     *client.Config
}

var (
	core *Core     //nolint:gochecknoglobals // pattern singleton
	once sync.Once //nolint:gochecknoglobals // pattern singleton
)

type OptFunc func(GophKeeperClient)

func (c *Core) SetStorage(r storage.Repo) {
	c.storage = r
}

func (c *Core) SetAPIClient(client api.Client) {
	c.client = client
}

func (c *Core) SetConfig(cfg *client.Config) {
	c.cfg = cfg
}

func (c *Core) InitDB() {
	c.storage.MigrateDB()
}

var (
	errPasswordCheck = errors.New("invalid password")
	errToken         = errors.New("invalid token")
)

func (c *Core) Sync(userPassword string) {
	if !c.verifyPassword(userPassword) {
		return
	}
	token, err := c.storage.GetSavedAccessToken()
	if err != nil {
		color.Red("Internal error: %v", err)

		return
	}
	c.loadCards(token)
	//c.loadLogins(token)
	//c.loadNotes(token)
}

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
