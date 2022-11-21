package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"
)

const createUserQuery = `
	INSERT INTO users (id, name, password)
	VALUES ($1, $2, $3)
`

func (p Storage) AddUser(ctx context.Context, name, hashedPassword string) (models.User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return models.User{}, err
	}
	_, err = p.conn.Exec(ctx, createUserQuery, id, name, hashedPassword)
	if err != nil {
		return models.User{}, err
	}
	return p.GetUserByName(ctx, name, hashedPassword)
}

const getUserByNameQuery = `
	SELECT * FROM users WHERE name = $1
`

func (p Storage) GetUserByName(ctx context.Context, name, hashedPassword string) (models.User, error) {
	var user models.User
	row, err := p.conn.Query(ctx, getUserByNameQuery, name)
	if err != nil {
		if err != pgx.ErrNoRows {
			return user, err
		}
		return user, nil
	}

	err = helpers.VerifyPassword(hashedPassword, user.Password)
	if err != nil {
		return user, err
	}

	err = row.Scan(user.ID, user.Name, user.Password)
	if err != nil {
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
