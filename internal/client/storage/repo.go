package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

type Repo interface {
	MigrateDB(ctx context.Context) error

	UserRepo
	CardStorage
	CredRepo
	NotesRepo
}

type UserRepo interface {
	AddUser(ctx context.Context, name string, password string) error
	UpdateUserToken(ctx context.Context, user *models.User, token *models.JWT) error
	DropUserToken(ctx context.Context, user *models.User) error
	GetAccessToken(ctx context.Context, user *models.User) (string, error)
	UserExists(ctx context.Context, name string) (bool, error)
	GetUserPasswordHash(ctx context.Context, user *models.User) (string, error)
}

type CardStorage interface {
	StoreCard(*models.Card) error
	StoreCards([]models.Card) error
	LoadCards() []models.Card
	GetCardByID(cardID uuid.UUID) (models.Card, error)
	DelCard(cardID uuid.UUID) error
}

type CredRepo interface {
	AddCred(*models.Cred) error
	SaveCreds([]models.Cred) error
	LoadCreds() []models.Cred
	GetCredByID(loginID uuid.UUID) (models.Cred, error)
	DelCred(loginID uuid.UUID) error
}

type NotesRepo interface {
	LoadNotes() []models.Note
	SaveNotes([]models.Note) error
	AddNote(*models.Note) error
	GetNoteByID(notedID uuid.UUID) (models.Note, error)
	DelNote(noteID uuid.UUID) error
}
