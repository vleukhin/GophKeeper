package core

import (
	"context"
	"log"

	"github.com/fatih/color"

	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *Core) Login(user *models.User) {
	ctx := context.TODO()
	token, err := c.client.Login(user)
	if err != nil {
		return
	}

	exists, err := c.storage.UserExists(ctx, user.Name)
	if err != nil {
		log.Fatal(err.Error())
	}

	if !exists {
		err = c.storage.AddUser(ctx, user.Name, user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = c.storage.UpdateUserToken(ctx, user, &token); err != nil {
		log.Fatal(err)
	}
	color.Green("Got authorization token for %q", user.Name)
}

func (c *Core) Register(user *models.User) {
	if err := c.client.Register(user); err != nil {
		return
	}
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = hashedPassword
	if err = c.storage.AddUser(context.TODO(), "", ""); err != nil {
		color.Red("Internal error: %v", err)

		return
	}

	color.Green("User registered")
	color.Green("ID: %v", user.ID)
	color.Green("Name: %s", user.Name)
}

func (c *Core) Logout() {
	if err := c.storage.DropUserToken(context.TODO(), nil); err != nil {
		color.Red("Internal error: %v", err)

		return
	}

	color.Green("Users tokens were dropped")
}

func (c *Core) verifyPassword(userPassword string) bool {
	hash, err := c.storage.GetUserPasswordHash(context.TODO(), nil)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if err := helpers.VerifyPassword(hash, userPassword); err != nil {
		color.Red("Password check status: failed")

		return false
	}

	return true
}
