package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (p Storage) LoadFiles(ctx context.Context) ([]models.File, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) SaveFiles(ctx context.Context, files []models.File) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) AddFile(ctx context.Context, file models.File) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) GetFileByID(ctx context.Context, fileID uuid.UUID) (models.File, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) DelFile(ctx context.Context, fileID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
