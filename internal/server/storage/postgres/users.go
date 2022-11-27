package postgres

import (
	"context"

	"github.com/vleukhin/GophKeeper/internal/helpers"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/vleukhin/GophKeeper/internal/models"
)

const createUserQuery = `
	INSERT INTO users (id, name, password)
	VALUES ($1, $2, $3)
`

func (p Storage) AddUser(ctx context.Context, name, password string) (models.User, error) {
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return models.User{}, err
	}
	_, err = p.conn.Exec(ctx, createUserQuery, id, name, hashedPassword)
	if err != nil {
		return models.User{}, err
	}
	return p.GetUserByName(ctx, name)
}

const getUserByNameQuery = `
	SELECT id, name, password FROM users WHERE name = $1
`

func (p Storage) GetUserByName(ctx context.Context, name string) (models.User, error) {
	var user models.User
	row := p.conn.QueryRow(ctx, getUserByNameQuery, name)
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil && err != pgx.ErrNoRows {
		return user, err
	}

	return user, nil
}

const getUserByIDQuery = `
	SELECT * FROM users WHERE id = $1
`

func (p Storage) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	row, err := p.conn.Query(ctx, getUserByIDQuery, id)
	if err != nil {
		if err != pgx.ErrNoRows {
			return user, err
		}
		return user, nil
	}

	err = row.Scan(user.ID, user.Name, user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
