package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (p Storage) GetCards(ctx context.Context, user models.User) ([]models.Card, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) AddCard(ctx context.Context, card *models.Card, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) UpdateCard(ctx context.Context, card *models.Card, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) IsCardOwner(ctx context.Context, cardUUID, userID uuid.UUID) bool {
	//TODO implement me
	panic("implement me")
}
