package client

import "github.com/vleukhin/GophKeeper/internal/client/auth"

type Core interface {
	auth.UserAuth
}

type App struct {
	*auth.Service
}

func NewApp() *App {
	return &App{
		auth.NewAuthService(),
	}
}
