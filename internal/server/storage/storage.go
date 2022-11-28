package storage

import (
	"context"

	"github.com/google/uuid"

	"github.com/vleukhin/GophKeeper/internal/models"
)

type (
	Repo interface {
		Ping(context.Context) error
		Migrate(ctx context.Context) error
		UsersStorage
		CredsStorage
		CardsStorage
		NotesStorage
		FileStorage
	}

	UsersStorage interface {
		AddUser(ctx context.Context, email, password string) (models.User, error)
		GetUserByName(ctx context.Context, name string) (models.User, error)
		GetUserByID(ctx context.Context, id string) (models.User, error)
	}

	CredsStorage interface {
		GetCreds(ctx context.Context, user models.User) ([]models.Cred, error)
		AddCred(ctx context.Context, login *models.Cred, userID uuid.UUID) error
		DelCred(ctx context.Context, loginID, userID uuid.UUID) error
		UpdateCred(ctx context.Context, login *models.Cred, userID uuid.UUID) error
		IsCredOwner(ctx context.Context, loginID, userID uuid.UUID) bool
	}

	CardsStorage interface {
		GetCards(ctx context.Context, user models.User) ([]models.Card, error)
		AddCard(ctx context.Context, card *models.Card, userID uuid.UUID) error
		DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error
		UpdateCard(ctx context.Context, card *models.Card, userID uuid.UUID) error
		IsCardOwner(ctx context.Context, cardUUID, userID uuid.UUID) bool
	}

	NotesStorage interface {
		GetNotes(ctx context.Context, user models.User) ([]models.Note, error)
		AddNote(ctx context.Context, note *models.Note, userID uuid.UUID) error
		DelNote(ctx context.Context, noteID, userID uuid.UUID) error
		UpdateNote(ctx context.Context, note *models.Note, userID uuid.UUID) error
		IsNoteOwner(ctx context.Context, noteID, userID uuid.UUID) bool
	}

	FileStorage interface {
		GetFiles(ctx context.Context, user models.User) ([]models.File, error)
		AddFile(ctx context.Context, binary models.File, userID uuid.UUID) error
		GetFile(ctx context.Context, binaryID, userID uuid.UUID) (models.File, error)
		DelFile(ctx context.Context, user models.User, binaryUUID uuid.UUID) error
	}
)
