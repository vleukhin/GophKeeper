package core

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
)

const minutesPerHour = 60

func (c *Core) SignUpUser(ctx context.Context, name, password string) (models.User, models.JWT, error) {
	user, err := c.repo.GetUserByName(ctx, name)
	if err != nil {
		return user, models.JWT{}, err
	}
	if !helpers.IsEmptyUUID(user.ID) {
		return user, models.JWT{}, errors.New("user already exists")
	}
	user, err = c.repo.AddUser(ctx, name, password)
	token, err := c.createToken(user)
	if err != nil {
		return user, token, err

	}
	return user, token, nil
}

func (c *Core) SignInUser(ctx context.Context, email, password string) (models.JWT, error) {
	user, err := c.repo.GetUserByName(ctx, email)
	if err != nil {
		return models.JWT{}, err
	}
	if helpers.IsEmptyUUID(user.ID) {
		return models.JWT{}, errors.New("user not found")
	}

	return c.createToken(user)
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
		log.Error().Err(err)
		return user, errs.ErrTokenValidation
	}

	return user, nil
}

func (c *Core) createToken(user models.User) (models.JWT, error) {
	var token models.JWT
	var err error

	token.AccessToken, err = helpers.CreateToken(
		c.cfg.Security.AccessTokenExpiresIn,
		user.ID,
		c.cfg.Security.AccessTokenPrivateKey)
	if err != nil {
		return token, err
	}

	token.RefreshToken, err = helpers.CreateToken(
		c.cfg.Security.RefreshTokenExpiresIn,
		user.ID,
		c.cfg.Security.RefreshTokenPrivateKey)

	if err != nil {
		return token, err
	}

	token.AccessTokenMaxAge = c.cfg.Security.AccessTokenMaxAge
	token.RefreshTokenMaxAge = c.cfg.Security.RefreshTokenMaxAge
	token.Domain = c.cfg.Security.Domain

	return token, nil
}
