package core

import (
	"context"
	"github.com/google/uuid"
	config "github.com/vleukhin/GophKeeper/internal/config/server"
	"github.com/vleukhin/GophKeeper/internal/models"
	"github.com/vleukhin/GophKeeper/internal/pkg/logger"
	"github.com/vleukhin/GophKeeper/internal/server/storage"
)

const minutesPerHour = 60

type GophKeeperServer interface {
	HealthCheck() error
	SignUpUser(ctx context.Context, email, password string) (models.User, error)
	SignInUser(ctx context.Context, email, password string) (models.JWT, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (models.JWT, error)
	GetDomainName() string
	CheckAccessToken(ctx context.Context, accessToken string) (models.User, error)

	GetLogins(ctx context.Context, user models.User) ([]models.Cred, error)
	AddLogin(ctx context.Context, login *models.Cred, userID uuid.UUID) error
	DelLogin(ctx context.Context, loginID, userID uuid.UUID) error
	UpdateLogin(ctx context.Context, login *models.Cred, userID uuid.UUID) error

	GetCards(ctx context.Context, user models.User) ([]models.Card, error)
	AddCard(ctx context.Context, card *models.Card, userID uuid.UUID) error
	DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error
	UpdateCard(ctx context.Context, card *models.Card, userID uuid.UUID) error

	GetNotes(ctx context.Context, user models.User) ([]models.Note, error)
	AddNote(ctx context.Context, note *models.Note, userID uuid.UUID) error
	DelNote(ctx context.Context, noteID, userID uuid.UUID) error
	UpdateNote(ctx context.Context, note *models.Note, userID uuid.UUID) error
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

func (c *Core) HealthCheck() error {
	return c.repo.Ping()
}

func (c *Core) GetDomainName() string {
	return c.cfg.Security.Domain
}