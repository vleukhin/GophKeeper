package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vleukhin/GophKeeper/internal/models"
	"time"
)

type Storage struct {
	conn *pgxpool.Pool
}

func NewPostgresStorage(dsn string, connTimeout time.Duration) (*Storage, error) {
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

func (p Storage) MigrateDB(ctx context.Context) error {
	migrations := []string{
		createUsersTableQuery,
		createCredsTableQuery,
		createCardsTableQuery,
		createNotesTableQuery,
	}
	for _, m := range migrations {
		_, err := p.conn.Query(ctx, m)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p Storage) AddUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) UpdateUserToken(user *models.User, token *models.JWT) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) DropUserToken() error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) GetSavedAccessToken() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) RemoveUsers() {
	//TODO implement me
	panic("implement me")
}

func (p Storage) UserExists(name string) bool {
	//TODO implement me
	panic("implement me")
}

func (p Storage) GetUserPasswordHash() string {
	//TODO implement me
	panic("implement me")
}

func (p Storage) StoreCard(card *models.Card) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) StoreCards(cards []models.Card) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) LoadCards() []models.Card {
	//TODO implement me
	panic("implement me")
}

func (p Storage) GetCardByID(cardID uuid.UUID) (models.Card, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) DelCard(cardID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) AddCred(cred *models.Cred) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) SaveCreds(creds []models.Cred) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) LoadCreds() []models.Cred {
	//TODO implement me
	panic("implement me")
}

func (p Storage) GetCredByID(loginID uuid.UUID) (models.Cred, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) DelCred(loginID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) LoadNotes() []models.Note {
	//TODO implement me
	panic("implement me")
}

func (p Storage) SaveNotes(notes []models.Note) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) AddNote(note *models.Note) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) GetNoteByID(notedID uuid.UUID) (models.Note, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) DelNote(noteID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
