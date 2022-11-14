package core

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"
	"log"
)

func (c *Core) loadLogins(accessToken string) {
	logins, err := c.client.GetCreds(accessToken)
	if err != nil {
		color.Red("Connection error: %v", err)

		return
	}

	if err = c.storage.SaveCreds(logins); err != nil {
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

	if err = c.storage.AddCred(login); err != nil {
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
	cred, err := c.storage.GetCredByID(loginUUID)
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

	if err := c.storage.DelCred(loginUUID); err != nil {
		log.Fatalf("Core - storage.DelLogin - %v", err)
	}

	if err := c.client.DelCred(accessToken, loginID); err != nil {
		log.Fatalf("Core - storage.DelLogin - %v", err)
	}

	color.Green("Login %q removed", loginID)
}