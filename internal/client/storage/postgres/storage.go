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

const storeCardQuery = `
	INSERT INTO cards (id, name, card_holder_name, number, bank, exp_month, exp_year, security_code, meta)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

func (p Storage) StoreCard(ctx context.Context, card models.Card) error {
	_, err := p.conn.Exec(ctx, storeCardQuery, card.ID, card.Name, card.CardHolderName, card.Number, card.Bank,
		card.ExpirationMonth, card.ExpirationYear, card.SecurityCode, card.Meta)

	return err
}

func (p Storage) StoreCards(ctx context.Context, cards []models.Card) error {
	var err error
	for _, c := range cards {
		err = p.StoreCard(ctx, c)
		if err != nil {
			return err
		}
	}

	return nil
}

const getCardByIDQuery = `
	SELECT * FROM cards where id = $1
`

func (p Storage) GetCardByID(ctx context.Context, cardID uuid.UUID) (models.Card, error) {
	var result models.Card
	row := p.conn.QueryRow(ctx, getCardByIDQuery, cardID)
	err := row.Scan(&result.ID, &result.Name, &result.CardHolderName, &result.Number, &result.Bank,
		&result.ExpirationMonth, &result.ExpirationYear, &result.SecurityCode, &result.Meta)
	if err != nil {
		if err == pgx.ErrNoRows {
			return result, nil
		}

		return result, err
	}

	return result, nil
}

const loadCardsQuery = `
	SELECT * from cards
`

func (p Storage) LoadCards(ctx context.Context) ([]models.Card, error) {
	var result []models.Card

	rows, err := p.conn.Query(ctx, loadCardsQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var card models.Card
		err := rows.Scan(&card.ID, &card.Name, &card.CardHolderName, &card.Number, &card.Bank,
			&card.ExpirationMonth, &card.ExpirationYear, &card.SecurityCode, &card.Meta)

		if err != nil {
			return nil, err
		}

		result = append(result, card)
	}

	return result, nil
}

const delCardQuery = `
	DELETE FROM cards WHERE id = $1
`

func (p Storage) DelCard(ctx context.Context, cardID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, delCardQuery, cardID)
	return err
}

const storeCredQuery = `
	INSERT INTO creds (id, name, login, password, meta)
	VALUES ($1, $2, $3, $4, $5) 
`

func (p Storage) AddCred(ctx context.Context, cred models.Cred) error {
	_, err := p.conn.Exec(ctx, storeCredQuery, cred.ID, cred.Name, cred.Login, cred.Password, cred.Meta)
	return err
}

func (p Storage) SaveCreds(ctx context.Context, creds []models.Cred) error {
	for _, c := range creds {
		err := p.AddCred(ctx, c)
		if err != nil {
			return err
		}
	}

	return nil
}

const loadCredsQuery = `
	SELECT * FROM creds
`

func (p Storage) LoadCreds(ctx context.Context) ([]models.Cred, error) {
	var result []models.Cred

	rows, err := p.conn.Query(ctx, loadCredsQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cred models.Cred
		err := rows.Scan(&cred.ID, &cred.Name, &cred.Login, &cred.Password, &cred.Meta)

		if err != nil {
			return nil, err
		}

		result = append(result, cred)
	}

	return result, nil
}

const getCredByID = `
	SELECT * FROM creds WHERE id = $1
`

func (p Storage) GetCredByID(ctx context.Context, loginID uuid.UUID) (models.Cred, error) {
	var result models.Cred
	row := p.conn.QueryRow(ctx, getCredByID, loginID)
	err := row.Scan(&result.ID, &result.Name, &result.Login, &result.Password, &result.Meta)
	if err != nil {
		if err == pgx.ErrNoRows {
			return result, nil
		}

		return result, err
	}

	return result, nil
}

const delCredQuery = `
	DELETE FROM creds WHERE id = $1
`

func (p Storage) DelCred(ctx context.Context, loginID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, delCredQuery, loginID)
	return err
}

const loadNotesQuery = `
	SELECT * FROM notes
`

func (p Storage) LoadNotes(ctx context.Context) ([]models.Note, error) {
	var result []models.Note

	rows, err := p.conn.Query(ctx, loadNotesQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Name, &note.Text, &note.Meta)

		if err != nil {
			return nil, err
		}

		result = append(result, note)
	}

	return result, nil
}

func (p Storage) SaveNotes(notes []models.Note) error {
	for _, n := range notes {
		err := p.AddNote(nil, n)
		if err != nil {
			return err
		}
	}

	return nil
}

const storeNoteQuery = `
	INSERT INTO notes (id, name, text, meta) VALUES ($1, $2, $3, $4)
`

func (p Storage) AddNote(ctx context.Context, note models.Note) error {
	_, err := p.conn.Exec(ctx, storeNoteQuery, note.ID, note.Text, note.Meta)
	return err
}

const getNoteByIDQuery = `
	SELECT * FROM notes WHERE id = $1
`

func (p Storage) GetNoteByID(ctx context.Context, notedID uuid.UUID) (models.Note, error) {
	var result models.Note
	row := p.conn.QueryRow(ctx, getNoteByIDQuery, notedID)
	err := row.Scan(&result.ID, &result.Name, &result.Text, &result.Meta)
	if err != nil {
		if err == pgx.ErrNoRows {
			return result, nil
		}

		return result, err
	}

	return result, nil
}

const delNoteQuery = `
	DELETE FROM notes WHERE id = $1
`

func (p Storage) DelNote(ctx context.Context, noteID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, delNoteQuery, noteID)
	return err
}
