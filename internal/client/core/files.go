package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *Core) StoreFile(file *os.File, name, filename string) {
	storedFile := models.File{
		ID:       uuid.UUID{},
		Name:     name,
		FileName: filename,
	}
	accessToken, err := c.authorisationCheck()
	if err != nil {
		return
	}
	err = c.client.StoreFile(accessToken, storedFile, helpers.EncryptStream(c.cfg.EncryptKey, file))
	if err != nil {
		color.Red(err.Error())
		return
	}
	err = c.storage.AddFile(context.Background(), storedFile)
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
	out, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		color.Red(err.Error())
		return
	}

	w := helpers.DecryptStream(c.cfg.EncryptKey, out)
	_, err = io.Copy(w, bytes.NewBuffer(file.Content))
	if err != nil {
		color.Red(err.Error())
		return
	}

	yellow := color.New(color.FgYellow).SprintFunc()
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
