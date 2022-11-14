package api

import "github.com/vleukhin/GophKeeper/internal/models"

const notesEndpoint = "api/notes"

func (c *HttpClient) GetNotes(accessToken string) (notes []models.Note, err error) {
	if err := c.getEntities(&notes, accessToken, notesEndpoint); err != nil {
		return nil, err
	}

	return notes, nil
}

func (c *HttpClient) StoreNote(accessToken string, note *models.Note) error {
	return c.addEntity(note, accessToken, notesEndpoint)
}

func (c *HttpClient) DelNote(accessToken, noteID string) error {
	return c.delEntity(accessToken, notesEndpoint, noteID)
}
