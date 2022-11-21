package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

type MockStorage struct {
}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func (m MockStorage) LoadNotes() []models.Note {
	return nil
}

func (m MockStorage) SaveNotes(notes []models.Note) error {
	return nil
}

func (m MockStorage) AddNote(note *models.Note) error {
	return nil
}

func (m MockStorage) GetNoteByID(notedID uuid.UUID) (models.Note, error) {
	return models.Note{}, nil
}

func (m MockStorage) DelNote(noteID uuid.UUID) error {
	return nil
}

func (m MockStorage) AddCred(cred *models.Cred) error {
	return nil
}

func (m MockStorage) SaveCreds(creds []models.Cred) error {
	return nil
}

func (m MockStorage) LoadCreds() []models.Cred {
	return nil
}

func (m MockStorage) GetCredByID(loginID uuid.UUID) (models.Cred, error) {
	return models.Cred{}, nil
}

func (m MockStorage) DelCred(loginID uuid.UUID) error {
	return nil
}

func (m MockStorage) MigrateDB(context.Context) error {
	return nil
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

func (m MockStorage) UserExists(name string) bool {
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
