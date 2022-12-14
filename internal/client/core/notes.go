package core

import (
	"context"
	"fmt"
	"log"

	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func (c *Core) loadNotes(accessToken string) {
	notes, err := c.client.GetNotes(accessToken)
	if err != nil {
		color.Red("Connection error: %v", err)

		return
	}

	if err = c.storage.SaveNotes(context.Background(), notes); err != nil {
		log.Println(err)

		return
	}
	color.Green("Loaded %v notes", len(notes))
}

func (c *Core) StoreNote(note *models.Note) {
	accessToken, err := c.authorisationCheck()
	if err != nil {
		log.Fatalf("Core - AddNote - %v", err)
	}
	c.encryptNote(c.cfg.EncryptKey, note)

	if err = c.client.StoreNote(accessToken, note); err != nil {
		log.Fatalf("Core - AddNote - %v", err)
	}

	if err = c.storage.AddNote(context.Background(), *note); err != nil {
		log.Fatal(err)
	}

	color.Green("Text %q added, id: %v", note.Name, note.ID)
}

func (c *Core) ShowNote(noteID string) {
	noteUUID, err := uuid.Parse(noteID)
	if err != nil {
		color.Red(err.Error())

		return
	}
	note, err := c.storage.GetNoteByID(context.Background(), noteUUID)
	if err != nil {
		color.Red(err.Error())

		return
	}

	c.decryptNote(c.cfg.EncryptKey, &note)
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("ID: %s\nname:%s\nText:%s\n%v\n", //nolint:forbidigo // cli printing
		yellow(note.ID),
		yellow(note.Name),
		yellow(note.Text),
		yellow(note.Meta),
	)
}

func (c *Core) encryptNote(key string, note *models.Note) {
	note.Text = helpers.Encrypt(key, note.Text)
}

func (c *Core) decryptNote(key string, note *models.Note) {
	note.Text = helpers.Decrypt(key, note.Text)
}

func (c *Core) DelNote(noteID string) {
	accessToken, err := c.authorisationCheck()
	if err != nil {
		return
	}
	noteUUID, err := uuid.Parse(noteID)
	if err != nil {
		color.Red(err.Error())
		log.Fatalf("Core - uuid.Parse - %v", err)
	}

	if err := c.storage.DelNote(context.Background(), noteUUID); err != nil {
		log.Fatalf("Core - storage.DelNote - %v", err)
	}

	if err := c.client.DelNote(accessToken, noteID); err != nil {
		log.Fatalf("Core - storage.DelNote - %v", err)
	}

	color.Green("Text %q removed", noteID)
}
