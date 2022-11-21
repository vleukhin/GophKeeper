package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vleukhin/GophKeeper/internal/models"
	"github.com/vleukhin/GophKeeper/internal/server/storage"
	"time"
)

type Storage struct {
	conn *pgxpool.Pool
}

func NewPostgresStorage(dsn string, connTimeout time.Duration) (storage.Repo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	conn, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Storage{
		conn: conn,
	}, nil
}

func (s Storage) Migrate(ctx context.Context) error {
	return nil
}

func (s Storage) Ping() error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) AddUser(ctx context.Context, email, hashedPassword string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetUserByEmail(ctx context.Context, email, hashedPassword string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetUserByID(ctx context.Context, id string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetLogins(ctx context.Context, user models.User) ([]models.Cred, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) AddLogin(ctx context.Context, login *models.Cred, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) DelLogin(ctx context.Context, loginID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) UpdateLogin(ctx context.Context, login *models.Cred, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) IsLoginOwner(ctx context.Context, loginID, userID uuid.UUID) bool {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetCards(ctx context.Context, user models.User) ([]models.Card, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) AddCard(ctx context.Context, card *models.Card, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) UpdateCard(ctx context.Context, card *models.Card, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) IsCardOwner(ctx context.Context, cardUUID, userID uuid.UUID) bool {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetNotes(ctx context.Context, user models.User) ([]models.Note, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) AddNote(ctx context.Context, note *models.Note, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) DelNote(ctx context.Context, noteID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
func (s Storage) UpdateNote(ctx context.Context, note *models.Note, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) IsNoteOwner(ctx context.Context, noteID, userID uuid.UUID) bool {
	//TODO implement me
	panic("implement me")
}
