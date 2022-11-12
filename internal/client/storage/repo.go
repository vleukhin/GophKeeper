package storage

import (
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

type Repo interface {
	MigrateDB()

	GetUserPasswordHash() string
	GetSavedAccessToken() (string, error)

	StoreCard(*models.Card) error
	StoreCards([]models.Card) error
	GetCardByID(cardID uuid.UUID) (models.Card, error)
	DelCard(cardID uuid.UUID) error
}
