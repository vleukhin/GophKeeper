package core

import (
	"context"

	"github.com/fatih/color"

	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *Core) Login(user models.User) {
	ctx := context.Background()
	token, err := c.client.Login(user)
	if err != nil {
		color.Red("Failed to log in user: %v", err)
		return
	}

	exists, err := c.storage.UserExists(ctx, user.Name)
	if err != nil {
		color.Red("Storage error: %v", err)
		return
	}

	if !exists {
		err = c.storage.AddUser(ctx, user.Name, user.Password)
		if err != nil {
			color.Red("Storage error: %v", err)
			return
		}
	}
	if err = c.storage.UpdateUserToken(ctx, user, token); err != nil {
		color.Red("Storage error: %v", err)
		return
	}
	color.Green("Got authorization token for %q", user.Name)
}

func (c *Core) Register(user models.User) {
	if err := c.client.Register(user); err != nil {
		color.Red("Failed to register user: %v", err)
		return
	}
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		color.Red("Failed to hash password: %v", err)
		return
	}

	user.Password = hashedPassword
	if err = c.storage.AddUser(context.Background(), user.Name, hashedPassword); err != nil {
		color.Red("Internal error: %v", err)
		return
	}

	color.Green("User registered")
	color.Green("ID: %v", user.ID)
	color.Green("Name: %s", user.Name)
}

func (c *Core) Logout() {
	if err := c.storage.DropUserToken(context.Background()); err != nil {
		color.Red("Storage error: %v", err)
		return
	}

	color.Green("Users tokens were dropped")
}

func (c *Core) verifyPassword(userPassword string) bool {
	hash, err := c.storage.GetUserPasswordHash(context.Background())
	if err != nil {
		color.Red("Storage error: %v", err)
		return false
	}
	if err := helpers.VerifyPassword(hash, userPassword); err != nil {
		color.Red("Password check status: failed")
		return false
	}

	return true
}
