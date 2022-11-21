package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func (p Storage) GetCreds(ctx context.Context, user models.User) ([]models.Cred, error) {
	//TODO implement me
	panic("implement me")
}

func (p Storage) AddCred(ctx context.Context, login *models.Cred, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) DelCred(ctx context.Context, loginID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) UpdateCred(ctx context.Context, login *models.Cred, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Storage) IsCredOwner(ctx context.Context, loginID, userID uuid.UUID) bool {
	//TODO implement me
	panic("implement me")
}
