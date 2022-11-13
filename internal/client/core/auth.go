package core

import (
	"log"

	"github.com/fatih/color"

	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *Core) Login(user *models.User) {
	token, err := c.client.Login(user)
	if err != nil {
		return
	}

	if !c.storage.UserExistsByEmail(user.Email) {
		err = c.storage.AddUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = c.storage.UpdateUserToken(user, &token); err != nil {
		log.Fatal(err)
	}
	color.Green("Got authorization token for %q", user.Email)
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
	if err = c.storage.AddUser(user); err != nil {
		color.Red("Internal error: %v", err)

		return
	}

	color.Green("User registered")
	color.Green("ID: %v", user.ID)
	color.Green("Email: %s", user.Email)
}

func (c *Core) Logout() {
	if err := c.storage.DropUserToken(); err != nil {
		color.Red("Internal error: %v", err)

		return
	}

	color.Green("Users tokens were dropped")
}

func (c *Core) verifyPassword(userPassword string) bool {
	if err := helpers.VerifyPassword(c.storage.GetUserPasswordHash(), userPassword); err != nil {
		color.Red("Password check status: failed")

		return false
	}

	return true
}
