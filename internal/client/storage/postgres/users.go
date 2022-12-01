package postgres

import (
	"context"

	"github.com/google/uuid"

	"github.com/vleukhin/GophKeeper/internal/models"
)

const createUserQuery = `
	INSERT INTO users (id, name, access_token, refresh_token)
	VALUES ($1, $2, $3, $4)
`

func (p Storage) AddUser(ctx context.Context, id uuid.UUID, name string, token models.JWT) error {
	_, err := p.conn.Exec(ctx, createUserQuery, id, name, token.AccessToken, token.RefreshToken)
	if err != nil {
		return err
	}

	return nil
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
