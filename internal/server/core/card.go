package core

import (
	"context"

	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/google/uuid"
)

func (c *Core) GetCards(ctx context.Context, user models.User) ([]models.Card, error) {
	return c.repo.GetCards(ctx, user)
}

func (c *Core) AddCard(ctx context.Context, card *models.Card, userID uuid.UUID) error {
	return c.repo.AddCard(ctx, card, userID)
}

func (c *Core) DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error {
	return c.repo.DelCard(ctx, cardUUID, userID)
}

func (c *Core) UpdateCard(ctx context.Context, card *models.Card, userID uuid.UUID) error {
	return c.repo.UpdateCard(ctx, card, userID)
}
