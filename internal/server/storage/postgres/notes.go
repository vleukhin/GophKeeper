package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (p Storage) GetNotes(ctx context.Context, user models.User) ([]models.Note, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) AddNote(ctx context.Context, note *models.Note, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) DelNote(ctx context.Context, noteID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) UpdateNote(ctx context.Context, note *models.Note, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) IsNoteOwner(ctx context.Context, noteID, userID uuid.UUID) bool {
	//TODO implement me
	panic("implement me")
}
