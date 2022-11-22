package postgres

import (
	"context"

	"github.com/google/uuid"

	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
	"github.com/vleukhin/GophKeeper/internal/models"
)

const getNotesQuery = `
	SELECT id, name, text, meta from notes where user_id = $1
`

func (p Storage) GetNotes(ctx context.Context, user models.User) ([]models.Note, error) {
	var notes []models.Note
	rows, err := p.conn.Query(ctx, getNotesQuery, user.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c models.Note
		err := rows.Scan(&c.ID, &c.Name, &c.Text, &c.Meta)
		if err != nil {
			return nil, err
		}

		notes = append(notes, c)
	}

	return notes, nil
}

const createNoteQuery = `
	INSERT INTO notes (id, user_id, name, text, meta)
	VALUES ($1, $2, $3, $4, $5) 
`

func (p Storage) AddNote(ctx context.Context, note *models.Note, userID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, createNoteQuery, note.ID, userID, note.Name, note.Text, note.Meta)
	return err
}

const delNoteQuery = `
	DELETE FROM notes WHERE id = $1
`

func (p Storage) DelNote(ctx context.Context, noteID, userID uuid.UUID) error {
	if !p.IsCardOwner(ctx, noteID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}
	_, err := p.conn.Exec(ctx, delNoteQuery, noteID)
	return err
}

const updateNoteQuery = `
	UPDATE notes SET
		 name = $1, 
		 text = $2,
		 meta = $4
	WHERE id = $5
`

func (p Storage) UpdateNote(ctx context.Context, note *models.Note, userID uuid.UUID) error {
	if !p.IsNoteOwner(ctx, note.ID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}
	_, err := p.conn.Exec(ctx, updateNoteQuery,
		note.Name,
		note.Text,
		note.Meta,
		userID,
	)

	return err
}

const getNoteByID = `
	SELECT id FROM notes WHERE id = $1 and user_id = $2
`

func (p Storage) IsNoteOwner(ctx context.Context, noteID, userID uuid.UUID) bool {
	return p.conn.QueryRow(ctx, getNoteByID, noteID, userID).Scan() == nil
}
