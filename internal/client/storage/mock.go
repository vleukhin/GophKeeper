package storage

import (
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

type MockStorage struct {
}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func (m MockStorage) MigrateDB() {

}

func (m MockStorage) AddUser(user *models.User) error {
	return nil
}

func (m MockStorage) UpdateUserToken(user *models.User, token *models.JWT) error {
	return nil
}

func (m MockStorage) DropUserToken() error {
	return nil
}

func (m MockStorage) GetSavedAccessToken() (string, error) {
	return "", nil
}

func (m MockStorage) RemoveUsers() {

}

func (m MockStorage) UserExistsByEmail(email string) bool {
	return false
}

func (m MockStorage) GetUserPasswordHash() string {
	return ""
}

func (m MockStorage) StoreCard(card *models.Card) error {
	return nil
}

func (m MockStorage) StoreCards(cards []models.Card) error {
	return nil
}

func (m MockStorage) LoadCards() []models.Card {
	return nil
}

func (m MockStorage) GetCardByID(cardID uuid.UUID) (models.Card, error) {
	return models.Card{}, nil
}

func (m MockStorage) DelCard(cardID uuid.UUID) error {
	return nil
}
