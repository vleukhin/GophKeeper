package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/vleukhin/GophKeeper/internal/models"
)

const loadNotesQuery = `
	SELECT * FROM notes
`

func (p Storage) LoadNotes(ctx context.Context) ([]models.Note, error) {
	var result []models.Note

	rows, err := p.conn.Query(ctx, loadNotesQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Name, &note.Text, &note.Meta)

		if err != nil {
			return nil, err
		}

		result = append(result, note)
	}

	return result, nil
}

func (p Storage) SaveNotes(ctx context.Context, notes []models.Note) error {
	for _, n := range notes {
		err := p.AddNote(ctx, n)
		if err != nil {
			return err
		}
	}

	return nil
}

const storeNoteQuery = `
	INSERT INTO notes (id, name, text, meta) VALUES ($1, $2, $3, $4)
`

func (p Storage) AddNote(ctx context.Context, note models.Note) error {
	_, err := p.conn.Exec(ctx, storeNoteQuery, note.ID, note.Text, note.Meta)
	return err
}

const getNoteByIDQuery = `
	SELECT * FROM notes WHERE id = $1
`

func (p Storage) GetNoteByID(ctx context.Context, notedID uuid.UUID) (models.Note, error) {
	var result models.Note
	row := p.conn.QueryRow(ctx, getNoteByIDQuery, notedID)
	err := row.Scan(&result.ID, &result.Name, &result.Text, &result.Meta)
	if err != nil {
		if err == pgx.ErrNoRows {
			return result, nil
		}

		return result, err
	}

	return result, nil
}

const delNoteQuery = `
	DELETE FROM notes WHERE id = $1
`

func (p Storage) DelNote(ctx context.Context, noteID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, delNoteQuery, noteID)
	return err
}
