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

func (m MockStorage) AddUser(ctx context.Context, name string, password string) error {
	return nil
}

func (m MockStorage) UpdateUserToken(ctx context.Context, user *models.User, token *models.JWT) error {
	return nil
}

func (m MockStorage) DropUserToken(context.Context, *models.User) error {
	return nil
}

func (m MockStorage) GetAccessToken(context.Context, *models.User) (string, error) {
	return "", nil
}

func (m MockStorage) UserExists(ctx context.Context, name string) (bool, error) {
	return false, nil
}

func (m MockStorage) GetUserPasswordHash(context.Context, *models.User) (string, error) {
	return "", nil
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
