package core

import (
	"context"
	"errors"
	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
)

const minutesPerHour = 60

func (c *Core) SignUpUser(ctx context.Context, name, password string) (user models.User, err error) {
	user, err = c.repo.GetUserByName(ctx, name)
	if err != nil {
		return
	}
	if user.Name != "" {
		return user, errors.New("user already exists")
	}
	return c.repo.AddUser(ctx, name, password)
}

func (c *Core) SignInUser(ctx context.Context, email, password string) (token models.JWT, err error) {
	user, err := c.repo.GetUserByName(ctx, email)
	if err != nil {
		return
	}

	err = helpers.VerifyPassword(user.Password, password)
	if err != nil {
		return
	}

	token.AccessToken, err = helpers.CreateToken(
		c.cfg.Security.AccessTokenExpiresIn,
		user.ID,
		c.cfg.Security.AccessTokenPrivateKey)
	if err != nil {
		return
	}

	token.RefreshToken, err = helpers.CreateToken(
		c.cfg.Security.RefreshTokenExpiresIn,
		user.ID,
		c.cfg.Security.RefreshTokenPrivateKey)

	if err != nil {
		return
	}

	token.AccessTokenMaxAge = c.cfg.Security.AccessTokenMaxAge
	token.RefreshTokenMaxAge = c.cfg.Security.RefreshTokenMaxAge
	token.Domain = c.cfg.Security.Domain

	return token, nil
}

func (c *Core) RefreshAccessToken(ctx context.Context, refreshToken string) (token models.JWT, err error) {
	userID, err := helpers.ValidateToken(refreshToken, c.cfg.Security.RefreshTokenPublicKey)
	if err != nil {
		err = errs.ErrTokenValidation

		return
	}

	user, err := c.repo.GetUserByID(ctx, userID)
	if err != nil {
		err = errs.ErrTokenValidation

		return
	}

	token.RefreshToken = refreshToken
	token.AccessToken, err = helpers.CreateToken(
		c.cfg.Security.AccessTokenExpiresIn,
		user.ID,
		c.cfg.Security.AccessTokenPrivateKey)

	if err != nil {
		return
	}

	token.AccessTokenMaxAge = c.cfg.Security.AccessTokenMaxAge * minutesPerHour
	token.RefreshTokenMaxAge = c.cfg.Security.RefreshTokenMaxAge * minutesPerHour
	token.Domain = c.cfg.Security.Domain

	return
}

func (c *Core) CheckAccessToken(ctx context.Context, accessToken string) (models.User, error) {
	var user models.User

	userID, err := helpers.ValidateToken(accessToken, c.cfg.AccessTokenPublicKey)
	if err != nil {
		err = errs.ErrTokenValidation

		return user, err
	}

	user, err = c.repo.GetUserByID(ctx, userID)

	if err != nil {
		return user, errs.ErrTokenValidation
	}

	return user, nil
}
