package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/vleukhin/GophKeeper/internal/models"
)

const loadFilesQuery = `
	SELECT * FROM files
`

func (p Storage) LoadFiles(ctx context.Context) ([]models.File, error) {
	var result []models.File

	rows, err := p.conn.Query(ctx, loadFilesQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var file models.File
		err := rows.Scan(&file.ID, &file.Name, &file.FileName, &file.Meta)

		if err != nil {
			return nil, err
		}

		result = append(result, file)
	}

	return result, nil
}

func (p Storage) SaveFiles(ctx context.Context, files []models.File) error {
	for _, n := range files {
		err := p.AddFile(ctx, n)
		if err != nil {
			return err
		}
	}

	return nil
}

const storeFileQuery = `
	INSERT INTO files (id, name, filename, meta) VALUES ($1, $2, $3, $4)
`

func (p Storage) AddFile(ctx context.Context, file models.File) error {
	_, err := p.conn.Exec(ctx, storeFileQuery, file.ID, file.Name, file.FileName, file.Meta)
	return err
}

const getFileByIDQuery = `
	SELECT * FROM files WHERE id = $1
`

func (p Storage) GetFileByID(ctx context.Context, filedID uuid.UUID) (models.File, error) {
	var result models.File
	row := p.conn.QueryRow(ctx, getFileByIDQuery, filedID)
	err := row.Scan(&result.ID, &result.Name, &result.FileName, &result.Meta)
	if err != nil {
		if err == pgx.ErrNoRows {
			return result, nil
		}

		return result, err
	}

	return result, nil
}

const delFileQuery = `
	DELETE FROM files WHERE id = $1
`

func (p Storage) DelFile(ctx context.Context, fileID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, delFileQuery, fileID)
	return err
}
