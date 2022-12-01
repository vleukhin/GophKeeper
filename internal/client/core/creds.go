package core

import (
	"context"
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

	if err = c.storage.SaveCreds(context.Background(), logins); err != nil {
		log.Println(err)

		return
	}
	color.Green("Loaded %v logins", len(logins))
}

func (c *Core) StoreCred(login *models.Cred) {
	accessToken, err := c.authorisationCheck()
	if err != nil {
		return
	}
	c.encryptLogin(c.cfg.EncryptKey, login)

	if err = c.client.AddCred(accessToken, login); err != nil {
		color.Red(err.Error())
		return
	}

	if err = c.storage.AddCred(context.Background(), *login); err != nil {
		color.Red(err.Error())
		return
	}

	color.Green("Login %q added, id: %v", login.Name, login.ID)
}

func (c *Core) ShowCred(loginID string) {
	loginUUID, err := uuid.Parse(loginID)
	if err != nil {
		color.Red(err.Error())
		return
	}
	cred, err := c.storage.GetCredByID(context.Background(), loginUUID)
	if err != nil {
		color.Red(err.Error())
		return
	}

	c.decryptLogin(c.cfg.EncryptKey, &cred)
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

func (c *Core) encryptLogin(key string, login *models.Cred) {
	login.Login = helpers.Encrypt(key, login.Login)
	login.Password = helpers.Encrypt(key, login.Password)
}

func (c *Core) decryptLogin(key string, login *models.Cred) {
	login.Login = helpers.Decrypt(key, login.Login)
	login.Password = helpers.Decrypt(key, login.Password)
}

func (c *Core) DelCred(loginID string) {
	accessToken, err := c.authorisationCheck()
	if err != nil {
		return
	}
	loginUUID, err := uuid.Parse(loginID)
	if err != nil {
		log.Fatalf("Core - uuid.Parse - %v", err)
	}

	if err := c.storage.DelCred(context.Background(), loginUUID); err != nil {
		log.Fatalf("Core - storage.DelCred - %v", err)
	}

	if err := c.client.DelCred(accessToken, loginID); err != nil {
		log.Fatalf("Core - storage.DelCred - %v", err)
	}

	color.Green("Login %q removed", loginID)
}
