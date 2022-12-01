package core

import (
	"context"
	"errors"
	"os"

	"github.com/fatih/color"

	"github.com/vleukhin/GophKeeper/internal/client/api"
	"github.com/vleukhin/GophKeeper/internal/client/storage"
	"github.com/vleukhin/GophKeeper/internal/config/client"
	"github.com/vleukhin/GophKeeper/internal/models"
)

type (
	GophKeeperClient interface {
		InitDB(ctx context.Context) error
		Sync()
		SetStorage(r storage.Repo)
		SetAPIClient(client api.Client)
		SetConfig(cfg *client.Config)

		AuthService
		CardsService
		NotesService
		CredService
		FilesService
	}

	AuthService interface {
		Register(name, password string)
		Login(name, password string)
		Logout()
	}
	CardsService interface {
		StoreCard(card *models.Card)
		ShowCard(cardID string)
		DelCard(cardID string)
	}
	NotesService interface {
		StoreNote(note *models.Note)
		ShowNote(noteID string)
		DelNote(noteID string)
	}
	CredService interface {
		StoreCred(login *models.Cred)
		ShowCred(loginID string)
		DelCred(loginID string)
	}
	FilesService interface {
		StoreFile(file *os.File, name, filename string)
		GetFile(fileID, filePath string)
		DelFile(fileID string)
	}
)

type Core struct {
	storage storage.Repo
	client  api.Client
	cfg     *client.Config
}

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

func (c *Core) InitDB(ctx context.Context) error {
	return c.storage.MigrateDB(ctx)
}

var (
	errToken = errors.New("invalid token")
)

func (c *Core) Sync() {
	token, err := c.storage.GetAccessToken(context.Background())
	if err != nil {
		color.Red("Internal error: %v", err)

		return
	}
	c.loadCards(token)
	c.loadLogins(token)
	c.loadNotes(token)
}

func (c *Core) authorisationCheck() (string, error) {
	accessToken, err := c.storage.GetAccessToken(context.Background())
	if err != nil || accessToken == "" {
		color.Red("User should be logged")

		return "", errToken
	}

	return accessToken, nil
}
