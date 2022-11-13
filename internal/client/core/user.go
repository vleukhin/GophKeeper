package core

import (
	"log"

	"github.com/fatih/color"

	"github.com/vleukhin/GophKeeper/internal/models"
	"github.com/vleukhin/GophKeeper/internal/utils"
)

func (c *Core) Login(user *models.User) {
	token, err := c.client.Login(user)
	if err != nil {
		return
	}

	if !c.storage.UserExistsByEmail(user.Email) {
		err = c.repo.AddUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = uc.repo.UpdateUserToken(user, &token); err != nil {
		log.Fatal(err)
	}
	color.Green("Got authorization token for %q", user.Email)
}

func (c *Core) Register(user *models.User) {
	if err := uc.clientAPI.Register(user); err != nil {
		return
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = hashedPassword
	if err = uc.repo.AddUser(user); err != nil {
		color.Red("Internal error: %v", err)

		return
	}

	color.Green("User registered")
	color.Green("ID: %v", user.ID)
	color.Green("Email: %s", user.Email)
}

func (c *Core) Logout() {
	if err := uc.repo.DropUserToken(); err != nil {
		color.Red("Internal error: %v", err)

		return
	}

	color.Green("Users tokens were dropped")
}

func (c *Core) Sync(userPassword string) {
	if !uc.verifyPassword(userPassword) {
		return
	}
	accessToken, err := uc.repo.GetSavedAccessToken()
	if err != nil {
		color.Red("Internal error: %v", err)

		return
	}
	uc.loadCards(accessToken)
	uc.loadLogins(accessToken)
	uc.loadNotes(accessToken)
}

func (c *Core) verifyPassword(userPassword string) bool {
	if err := utils.VerifyPassword(uc.repo.GetUserPasswordHash(), userPassword); err != nil {
		color.Red("Password check status: failed")

		return false
	}

	return true
}
