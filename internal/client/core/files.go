package core

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"
	"os"
)

func (c *Core) StoreFile(file models.File) {
	accessToken, err := c.authorisationCheck()
	if err != nil {
		return
	}

	c.encryptFile(c.cfg.EncryptKey, &file)
	err = c.client.StoreFile(accessToken, file)
	if err != nil {
		color.Red(err.Error())
		return
	}

	err = c.storage.AddFile(context.Background(), file)
	if err != nil {
		color.Red(err.Error())
		return
	}
}

func (c *Core) GetFile(fileID, filePath string) {
	fileUUID, err := uuid.Parse(fileID)
	if err != nil {
		color.Red(err.Error())
		return
	}
	file, err := c.storage.GetFileByID(context.Background(), fileUUID)
	if err != nil {
		color.Red(err.Error())
		return
	}
	c.decryptFile(c.cfg.EncryptKey, &file)
	yellow := color.New(color.FgYellow).SprintFunc()
	err = os.WriteFile(filePath, file.Content, 0600)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("File saved to %s", yellow(filePath))
}

func (c *Core) DelFile(fileID string) {
	accessToken, err := c.authorisationCheck()
	if err != nil {
		return
	}
	fileUUID, err := uuid.Parse(fileID)
	if err != nil {
		color.Red(err.Error())
		return
	}
	err = c.client.DelCred(accessToken, fileID)
	if err != nil {
		color.Red(err.Error())
		return
	}
	err = c.storage.DelFile(context.Background(), fileUUID)
	if err != nil {
		color.Red(err.Error())
		return
	}

	color.Green("Card %q removed", fileID)
}

func (c *Core) encryptFile(key string, file *models.File) {
	file.Content = []byte(helpers.Encrypt(key, string(file.Content)))
}

func (c *Core) decryptFile(key string, file *models.File) {
	file.Content = []byte(helpers.Decrypt(key, string(file.Content)))
}
