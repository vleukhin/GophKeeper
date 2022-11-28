package postgres

import (
	"context"

	"github.com/google/uuid"

	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
	"github.com/vleukhin/GophKeeper/internal/models"
)

const getFilesQuery = `
	SELECT id, name, meta
	FROM files where user_id = $1
`

func (p Storage) GetFiles(ctx context.Context, user models.User) ([]models.File, error) {
	var result []models.File

	rows, err := p.conn.Query(ctx, getFilesQuery, user.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var file models.File
		err := rows.Scan(&file.ID, &file.Name, &file.Meta)
		if err != nil {
			return nil, err
		}

		result = append(result, file)
	}

	return result, nil
}

const storeFileQuery = `
	INSERT INTO files (id, user_id, name, filename, meta)
	VALUES ($1, $2, $3, $4, $5)
`

func (p Storage) AddFile(ctx context.Context, file models.File, userID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, storeFileQuery, file.ID, userID, file.Name, file.FileName, file.Meta)
	return err
}

const delFileQuery = `
	DELETE FROM cards WHERE id = $1
`

func (p Storage) DelFile(ctx context.Context, user models.User, fileUUID uuid.UUID) error {
	if !p.IsCardOwner(ctx, fileUUID, user.ID) {
		return errs.ErrWrongOwnerOrNotFound
	}
	_, err := p.conn.Exec(ctx, delFileQuery, fileUUID)
	return err
}

const getFileByID = `
	SELECT id FROM files WHERE id = $1 and user_id = $2
`

func (p Storage) IsFileOwner(ctx context.Context, fileUUID, userID uuid.UUID) bool {
	return p.conn.QueryRow(ctx, getFileByID, fileUUID, userID).Scan() == nil
}
