package core

import (
	"context"

	"github.com/google/uuid"

	config "github.com/vleukhin/GophKeeper/internal/config/server"
	"github.com/vleukhin/GophKeeper/internal/models"
	"github.com/vleukhin/GophKeeper/internal/pkg/logger"
	"github.com/vleukhin/GophKeeper/internal/server/storage"
)

type GophKeeperServer interface {
	HealthCheck(context.Context) error
	SignUpUser(ctx context.Context, name, password string) (models.User, models.JWT, error)
	SignInUser(ctx context.Context, name, password string) (models.JWT, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (models.JWT, error)
	GetDomainName() string
	CheckAccessToken(ctx context.Context, accessToken string) (models.User, error)

	GetCred(ctx context.Context, user models.User) ([]models.Cred, error)
	AddCred(ctx context.Context, login *models.Cred, userID uuid.UUID) error
	DelCred(ctx context.Context, loginID, userID uuid.UUID) error
	UpdateCred(ctx context.Context, login *models.Cred, userID uuid.UUID) error

	GetCards(ctx context.Context, user models.User) ([]models.Card, error)
	AddCard(ctx context.Context, card *models.Card, userID uuid.UUID) error
	DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error
	UpdateCard(ctx context.Context, card *models.Card, userID uuid.UUID) error

	GetNotes(ctx context.Context, user models.User) ([]models.Note, error)
	AddNote(ctx context.Context, note *models.Note, userID uuid.UUID) error
	DelNote(ctx context.Context, noteID, userID uuid.UUID) error
	UpdateNote(ctx context.Context, note *models.Note, userID uuid.UUID) error

	GetFiles(ctx context.Context, user models.User) ([]models.File, error)
	AddFile(ctx context.Context, file models.File, userID uuid.UUID) error
	DelFile(ctx context.Context, user models.User, fileUUID uuid.UUID) error
}

type Core struct {
	repo storage.Repo
	cfg  *config.Config
	l    *logger.Logger
}

func New(r storage.Repo, cfg *config.Config, l *logger.Logger) *Core {
	return &Core{
		repo: r,
		cfg:  cfg,
		l:    l,
	}
}

func (c *Core) HealthCheck(ctx context.Context) error {
	return c.repo.Ping(ctx)
}

func (c *Core) GetDomainName() string {
	return c.cfg.Security.Domain
}
