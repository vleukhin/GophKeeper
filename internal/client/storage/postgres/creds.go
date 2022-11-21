package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/vleukhin/GophKeeper/internal/models"
)

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
