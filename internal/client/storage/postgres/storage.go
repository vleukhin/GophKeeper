package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vleukhin/GophKeeper/internal/helpers"
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

const createUserQuery = `
	INSERT INTO users (id, name, password)
	VALUES ($1, $2, $3)
`

func (p Storage) AddUser(ctx context.Context, name string, password string) error {
	hPassword, err := helpers.HashPassword(password)
	if err != nil {
		return err
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	_, err = p.conn.Exec(ctx, createUserQuery, id, name, hPassword)
	if err != nil {
		return err
	}
	return nil
}

const updateUserTokenQuery = `
	UPDATE users SET access_token = $1, refresh_token = $2
	WHERE id = $3
`

func (p Storage) UpdateUserToken(ctx context.Context, user *models.User, token *models.JWT) error {
	_, err := p.conn.Exec(ctx, updateUserTokenQuery, token.AccessToken, token.RefreshToken, user.ID)
	return err
}

const dropUserTokenQuery = `
	UPDATE users SET access_token = '', refresh_token = ''
	WHERE id = $3
`

func (p Storage) DropUserToken(ctx context.Context, user *models.User) error {
	_, err := p.conn.Exec(ctx, dropUserTokenQuery, user.ID)
	return err
}

const getAccessTokenQuery = `
	SELECT access_token FROM users WHERE id = $1
`

func (p Storage) GetAccessToken(ctx context.Context, user *models.User) (string, error) {
	var token string
	row := p.conn.QueryRow(ctx, getAccessTokenQuery, user.ID)
	err := row.Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}

const getUserQuery = `
	SELECT id FROM users WHERE name = $1
`

func (p Storage) UserExists(ctx context.Context, name string) (bool, error) {
	_, err := p.conn.Query(ctx, getUserQuery, name)
	if err != nil {
		if err != pgx.ErrNoRows {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

const getUserPasswordHashQuery = `
	SELECT password FROM users WHERE id = $1
`

func (p Storage) GetUserPasswordHash(ctx context.Context, user *models.User) (string, error) {
	var password string
	row := p.conn.QueryRow(ctx, getUserPasswordHashQuery, user.ID)
	err := row.Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
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
