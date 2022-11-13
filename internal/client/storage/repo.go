package storage

import (
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

type Repo interface {
	MigrateDB()

	UserRepo
	CardStorage
}

type UserRepo interface {
	AddUser(user *models.User) error
	UpdateUserToken(user *models.User, token *models.JWT) error
	DropUserToken() error
	GetSavedAccessToken() (string, error)
	RemoveUsers()
	UserExistsByEmail(email string) bool
	GetUserPasswordHash() string
}

type CardStorage interface {
	StoreCard(*models.Card) error
	StoreCards([]models.Card) error
	LoadCards() []models.Card
	GetCardByID(cardID uuid.UUID) (models.Card, error)
	DelCard(cardID uuid.UUID) error
}
