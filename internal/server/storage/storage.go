package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

type (
	Repo interface {
		Ping() error
		Migrate(ctx context.Context) error
		UsersStorage
		CredsStorage
		CardsStorage
		NotesStorage
	}

	UsersStorage interface {
		AddUser(ctx context.Context, email, hashedPassword string) (models.User, error)
		GetUserByEmail(ctx context.Context, email, hashedPassword string) (models.User, error)
		GetUserByID(ctx context.Context, id string) (models.User, error)
	}

	CredsStorage interface {
		GetLogins(ctx context.Context, user models.User) ([]models.Cred, error)
		AddLogin(ctx context.Context, login *models.Cred, userID uuid.UUID) error
		DelLogin(ctx context.Context, loginID, userID uuid.UUID) error
		UpdateLogin(ctx context.Context, login *models.Cred, userID uuid.UUID) error
		IsLoginOwner(ctx context.Context, loginID, userID uuid.UUID) bool
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
)