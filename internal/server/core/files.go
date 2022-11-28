package core

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
	"mime/multipart"
)

func (c *Core) GetFiles(ctx context.Context, user models.User) ([]models.File, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Core) AddFile(ctx context.Context, binary models.File, file *multipart.FileHeader, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c *Core) GetFile(ctx context.Context, currentUser models.User, binaryUUID uuid.UUID) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Core) DelFile(ctx context.Context, currentUser models.User, binaryUUID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
