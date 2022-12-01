package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/vleukhin/GophKeeper/internal/models"
)

const createUserQuery = `
	INSERT INTO users (id, name, access_token, refresh_token)
	VALUES ($1, $2, $3, $4)
`

func (p Storage) AddUser(ctx context.Context, name string, token models.JWT) (models.User, error) {
	var user models.User
	id, err := uuid.NewUUID()
	if err != nil {
		return user, err
	}
	_, err = p.conn.Exec(ctx, createUserQuery, id, name, token.AccessToken, token.RefreshToken)
	if err != nil {
		return user, err
	}
	user.ID = id
	user.Name = name

	return user, nil
}

const dropUserQuery = `
	DELETE FROM users
`

func (p Storage) DropUser(ctx context.Context) error {
	_, err := p.conn.Exec(ctx, dropUserQuery)
	return err
}

const getAccessTokenQuery = `
	SELECT access_token FROM users
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
