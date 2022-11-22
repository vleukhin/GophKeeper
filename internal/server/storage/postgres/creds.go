package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
	"github.com/vleukhin/GophKeeper/internal/models"
)

const getCredsQuery = `
	SELECT id, name, login, password, meta from creds where user_id = $1
`

func (p Storage) GetCreds(ctx context.Context, user models.User) ([]models.Cred, error) {
	var creds []models.Cred
	rows, err := p.conn.Query(ctx, getCredsQuery, user.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c models.Cred
		err := rows.Scan(&c.ID, &c.Name, &c.Login, &c.Password, &c.Meta)
		if err != nil {
			return nil, err
		}

		creds = append(creds, c)
	}

	return creds, nil
}

const createCredQuery = `
	INSERT INTO creds (id, user_id, name, login, password, meta)
	VALUES ($1, $2, $3, $4, $5, $6) 
`

func (p Storage) AddCred(ctx context.Context, cred *models.Cred, userID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, createCredQuery, cred.ID, userID, cred.Name, cred.Login, cred.Password, cred.Meta)
	return err
}

const delCredQuery = `
	DELETE FROM creds WHERE id = $1
`

func (p Storage) DelCred(ctx context.Context, credID, userID uuid.UUID) error {
	if !p.IsCardOwner(ctx, credID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}
	_, err := p.conn.Exec(ctx, delCredQuery, credID)
	return err
}

const updateCredQuery = `
	UPDATE creds SET
		 name = $1, 
		 login = $2,
		 password = $3, 
		 meta = $4
	WHERE id = $5
`

func (p Storage) UpdateCred(ctx context.Context, cred *models.Cred, userID uuid.UUID) error {
	if !p.IsCredOwner(ctx, cred.ID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}
	_, err := p.conn.Exec(ctx, updateCredQuery,
		cred.Name,
		cred.Login,
		cred.Password,
		cred.Meta,
		userID,
	)

	return err
}

const getCredByID = `
	SELECT id FROM creds WHERE id = $1 and user_id = $2
`

func (p Storage) IsCredOwner(ctx context.Context, credID, userID uuid.UUID) bool {
	return p.conn.QueryRow(ctx, getCredByID, credID, userID).Scan() == nil
}
