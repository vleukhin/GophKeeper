package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (p Storage) GetFiles(ctx context.Context, user models.User) ([]models.File, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) AddFile(ctx context.Context, binary models.File, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) GetFile(ctx context.Context, binaryID, userID uuid.UUID) (models.File, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) DelFile(ctx context.Context, user models.User, binaryUUID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
