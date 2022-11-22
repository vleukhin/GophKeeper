package core

import (
	"context"

	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/google/uuid"
)

func (c *Core) GetLogins(ctx context.Context, user models.User) ([]models.Cred, error) {
	return c.repo.GetCreds(ctx, user)
}

func (c *Core) AddLogin(ctx context.Context, login *models.Cred, userID uuid.UUID) error {
	return c.repo.AddCred(ctx, login, userID)
}

func (c *Core) DelLogin(ctx context.Context, loginID, userID uuid.UUID) error {
	return c.repo.DelCred(ctx, loginID, userID)
}

func (c *Core) UpdateLogin(ctx context.Context, login *models.Cred, userID uuid.UUID) error {
	return c.repo.UpdateCred(ctx, login, userID)
}
