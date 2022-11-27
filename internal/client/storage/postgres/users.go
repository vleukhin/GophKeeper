package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/vleukhin/GophKeeper/internal/models"
)

const createUserQuery = `
	INSERT INTO users (id, name, password)
	VALUES ($1, $2, $3)
`

func (p Storage) AddUser(ctx context.Context, name string, password string) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	_, err = p.conn.Exec(ctx, createUserQuery, id, name, password)
	if err != nil {
		return err
	}
	return nil
}

const updateUserTokenQuery = `
	UPDATE users SET access_token = $1, refresh_token = $2
	WHERE id = $3
`

func (p Storage) UpdateUserToken(ctx context.Context, user models.User, token models.JWT) error {
	_, err := p.conn.Exec(ctx, updateUserTokenQuery, token.AccessToken, token.RefreshToken, user.ID)
	return err
}

const dropUserTokenQuery = `
	UPDATE users SET access_token = '', refresh_token = ''
	WHERE id = $3
`

func (p Storage) DropUserToken(ctx context.Context) error {
	_, err := p.conn.Exec(ctx, dropUserTokenQuery)
	return err
}

const getAccessTokenQuery = `
	SELECT access_token FROM users WHERE id = $1
`

func (p Storage) GetAccessToken(ctx context.Context) (string, error) {
	var token string
	row := p.conn.QueryRow(ctx, getAccessTokenQuery)
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

func (p Storage) GetUserPasswordHash(ctx context.Context) (string, error) {
	var password string
	row := p.conn.QueryRow(ctx, getUserPasswordHashQuery)
	err := row.Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}
