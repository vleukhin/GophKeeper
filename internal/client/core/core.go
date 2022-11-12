package core

import (
	"errors"
	"github.com/fatih/color"
	"github.com/vleukhin/GophKeeper/internal/client/storage"
	"github.com/vleukhin/GophKeeper/internal/config/client"
	"github.com/vleukhin/GophKeeper/internal/helpers"
)

var (
	errPasswordCheck = errors.New("incorrect password")
	errToken         = errors.New("token error")
)

type App interface {
	storage.Repo
	storage.Client
}

type Core struct {
	repo   storage.Repo
	client storage.Client
	cfg    *client.Config
}

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

func (с *Core) verifyPassword(userPassword string) bool {
	if err := helpers.VerifyPassword(с.repo.GetUserPasswordHash(), userPassword); err != nil {
		color.Red("Password check status: failed")

		return false
	}

	return true
}
