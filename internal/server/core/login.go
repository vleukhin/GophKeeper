package core

import (
	"context"

	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/google/uuid"
)

func (c *Core) GetCred(ctx context.Context, user models.User) ([]models.Cred, error) {
	return c.repo.GetCreds(ctx, user)
}

func (c *Core) AddCred(ctx context.Context, login *models.Cred, userID uuid.UUID) error {
	return c.repo.AddCred(ctx, login, userID)
}

func (c *Core) DelCred(ctx context.Context, loginID, userID uuid.UUID) error {
	return c.repo.DelCred(ctx, loginID, userID)
}

func (c *Core) UpdateCred(ctx context.Context, login *models.Cred, userID uuid.UUID) error {
	return c.repo.UpdateCred(ctx, login, userID)
}
