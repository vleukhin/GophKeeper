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
	FilesRepo
}

type UserRepo interface {
	AddUser(ctx context.Context, name string, password string) (models.User, error)
	UpdateUserToken(ctx context.Context, user models.User, token models.JWT) error
	DropUser(ctx context.Context) error
	GetAccessToken(ctx context.Context) (string, error)
	UserExists(ctx context.Context, name string) (bool, error)
	GetUserPasswordHash(ctx context.Context) (string, error)
}

type CardStorage interface {
	StoreCard(context.Context, models.Card) error
	StoreCards(context.Context, []models.Card) error
	LoadCards(context.Context) ([]models.Card, error)
	GetCardByID(ctx context.Context, cardID uuid.UUID) (models.Card, error)
	DelCard(ctx context.Context, cardID uuid.UUID) error
}

type CredRepo interface {
	AddCred(context.Context, models.Cred) error
	SaveCreds(context.Context, []models.Cred) error
	LoadCreds(context.Context) ([]models.Cred, error)
	GetCredByID(ctx context.Context, loginID uuid.UUID) (models.Cred, error)
	DelCred(ctx context.Context, loginID uuid.UUID) error
}

type NotesRepo interface {
	LoadNotes(context.Context) ([]models.Note, error)
	SaveNotes(context.Context, []models.Note) error
	AddNote(context.Context, models.Note) error
	GetNoteByID(ctx context.Context, notedID uuid.UUID) (models.Note, error)
	DelNote(ctx context.Context, noteID uuid.UUID) error
}

type FilesRepo interface {
	LoadFiles(context.Context) ([]models.File, error)
	SaveFiles(context.Context, []models.File) error
	AddFile(context.Context, models.File) error
	GetFileByID(ctx context.Context, fileID uuid.UUID) (models.File, error)
	DelFile(ctx context.Context, fileID uuid.UUID) error
}
