package core

import (
	"context"

	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/google/uuid"
)

func (c *Core) GetNotes(ctx context.Context, user models.User) ([]models.Note, error) {
	return c.repo.GetNotes(ctx, user)
}

func (c *Core) AddNote(ctx context.Context, note *models.Note, userID uuid.UUID) error {
	return c.repo.AddNote(ctx, note, userID)
}

func (c *Core) DelNote(ctx context.Context, noteID, userID uuid.UUID) error {
	return c.repo.DelNote(ctx, noteID, userID)
}

func (c *Core) UpdateNote(ctx context.Context, note *models.Note, userID uuid.UUID) error {
	return c.repo.UpdateNote(ctx, note, userID)
}
