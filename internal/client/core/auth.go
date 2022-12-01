package core

import (
	"context"

	"github.com/fatih/color"
)

func (c *Core) Login(name, password string) {
	ctx := context.Background()
	user, token, err := c.client.Login(name, password)

	if err != nil {
		color.Red("Failed to log in user: %v", err)
		return
	}
	if err := c.storage.DropUser(ctx); err != nil {
		color.Red("Storage error: %v", err)
		return
	}
	err = c.storage.AddUser(ctx, user.ID, name, token)
	if err != nil {
		color.Red("Storage error: %v", err)
		return
	}

	color.Green("Got authorization token for %q", user.Name)
}

func (c *Core) Register(name, password string) {
	ctx := context.Background()
	user, token, err := c.client.Register(name, password)
	if err != nil {
		color.Red("Failed to register user: %v", err)
		return
	}
	if err := c.storage.DropUser(ctx); err != nil {
		color.Red("Storage error: %v", err)
		return
	}

	err = c.storage.AddUser(ctx, user.ID, user.Name, token)
	if err != nil {
		color.Red("Internal error: %v", err)
		return
	}

	color.Green("User registered")
	color.Green("ID: %v", user.ID.String())
	color.Green("Name: %s", user.Name)
}

func (c *Core) Logout() {
	if err := c.storage.DropUser(context.Background()); err != nil {
		color.Red("Storage error: %v", err)
		return
	}

	color.Green("Users tokens were dropped")
}
