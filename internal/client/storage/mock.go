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

func (m MockStorage) LoadNotes(context.Context) ([]models.Note, error) {
	return nil, nil
}

func (m MockStorage) SaveNotes(ctx context.Context, notes []models.Note) error {
	return nil
}

func (m MockStorage) AddNote(ctx context.Context, note models.Note) error {
	return nil
}

func (m MockStorage) GetNoteByID(ctx context.Context, notedID uuid.UUID) (models.Note, error) {
	return models.Note{}, nil
}

func (m MockStorage) DelNote(ctx context.Context, noteID uuid.UUID) error {
	return nil
}

func (m MockStorage) AddCred(ctx context.Context, cred models.Cred) error {
	return nil
}

func (m MockStorage) SaveCreds(ctx context.Context, creds []models.Cred) error {
	return nil
}

func (m MockStorage) LoadCreds(context.Context) ([]models.Cred, error) {
	return nil, nil
}

func (m MockStorage) GetCredByID(ctx context.Context, loginID uuid.UUID) (models.Cred, error) {
	return models.Cred{}, nil
}

func (m MockStorage) DelCred(ctx context.Context, loginID uuid.UUID) error {
	return nil
}

func (m MockStorage) MigrateDB(context.Context) error {
	return nil
}

func (m MockStorage) AddUser(ctx context.Context, name string, password string) error {
	return nil
}

func (m MockStorage) UpdateUserToken(ctx context.Context, user models.User, token models.JWT) error {
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

func (m MockStorage) StoreCard(ctx context.Context, card models.Card) error {
	return nil
}

func (m MockStorage) StoreCards(ctx context.Context, cards []models.Card) error {
	return nil
}

func (m MockStorage) LoadCards(context.Context) ([]models.Card, error) {
	return nil, nil
}

func (m MockStorage) GetCardByID(ctx context.Context, cardID uuid.UUID) (models.Card, error) {
	return models.Card{}, nil
}

func (m MockStorage) DelCard(ctx context.Context, cardID uuid.UUID) error {
	return nil
}
