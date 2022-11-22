package core

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *Core) loadLogins(accessToken string) {
	logins, err := c.client.GetCreds(accessToken)
	if err != nil {
		color.Red("Connection error: %v", err)

		return
	}

	if err = c.storage.SaveCreds(nil, logins); err != nil {
		log.Println(err)

		return
	}
	color.Green("Loaded %v logins", len(logins))
}

func (c *Core) StoreCred(userPassword string, login *models.Cred) {
	accessToken, err := c.authorisationCheck(userPassword)
	if err != nil {
		return
	}
	c.encryptLogin(userPassword, login)

	if err = c.client.AddCred(accessToken, login); err != nil {
		return
	}

	if err = c.storage.AddCred(nil, *login); err != nil {
		log.Fatal(err)
	}

	color.Green("Login %q added, id: %v", login.Name, login.ID)
}

func (c *Core) ShowCred(userPassword, loginID string) {
	if !c.verifyPassword(userPassword) {
		return
	}
	loginUUID, err := uuid.Parse(loginID)
	if err != nil {
		color.Red(err.Error())

		return
	}
	cred, err := c.storage.GetCredByID(nil, loginUUID)
	if err != nil {
		color.Red(err.Error())

		return
	}

	c.decryptLogin(userPassword, &cred)
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("ID: %s\nname:%s\nURI:%s\nLogin:%s\nPassword:%s\n%v\n",
		yellow(cred.ID),
		yellow(cred.Name),
		yellow(cred.URI),
		yellow(cred.Login),
		yellow(cred.Password),
		yellow(cred.Meta),
	)
}

func (c *Core) encryptLogin(userPassword string, login *models.Cred) {
	login.Login = helpers.Encrypt(userPassword, login.Login)
	login.Password = helpers.Encrypt(userPassword, login.Password)
}

func (c *Core) decryptLogin(userPassword string, login *models.Cred) {
	login.Login = helpers.Decrypt(userPassword, login.Login)
	login.Password = helpers.Decrypt(userPassword, login.Password)
}

func (c *Core) DelCred(userPassword, loginID string) {
	accessToken, err := c.authorisationCheck(userPassword)
	if err != nil {
		return
	}
	loginUUID, err := uuid.Parse(loginID)
	if err != nil {
		color.Red(err.Error())
		log.Fatalf("Core - uuid.Parse - %v", err)
	}

	if err := c.storage.DelCred(nil, loginUUID); err != nil {
		log.Fatalf("Core - storage.DelCred - %v", err)
	}

	if err := c.client.DelCred(accessToken, loginID); err != nil {
		log.Fatalf("Core - storage.DelCred - %v", err)
	}

	color.Green("Login %q removed", loginID)
}
