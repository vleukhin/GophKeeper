package core

import (
	"context"
	"fmt"
	"github.com/vleukhin/GophKeeper/internal/helpers"
	"os"

	"github.com/google/uuid"

	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *Core) GetFiles(ctx context.Context, user models.User) ([]models.File, error) {
	files, err := c.repo.GetFiles(ctx, user)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		content, err := os.ReadFile(helpers.FilePath(c.cfg.FilesStorage.Location, user.ID.String(), f.ID.String()))
		if err != nil {
			return nil, err
		}
		f.Content = content
	}

	return files, nil
}

func (c *Core) AddFile(ctx context.Context, file models.File, userID uuid.UUID) error {
	file.ID = uuid.New()
	fmt.Println(helpers.FileDir(c.cfg.FilesStorage.Location, userID.String()))
	fmt.Println(c.cfg.FilesStorage.Location)
	err := os.MkdirAll(helpers.FileDir(c.cfg.FilesStorage.Location, userID.String()), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(helpers.FilePath(c.cfg.FilesStorage.Location, userID.String(), file.ID.String()), file.Content, 0600)
	if err != nil {
		return err
	}
	return c.repo.AddFile(ctx, file, userID)
}

func (c *Core) DelFile(ctx context.Context, user models.User, fileID uuid.UUID) error {
	err := os.Remove(helpers.FilePath(c.cfg.FilesStorage.Location, user.ID.String(), fileID.String()))
	if err != nil {
		return err
	}
	return c.repo.DelFile(ctx, user, fileID)
}
