package core

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
)

const minutesPerHour = 60

func (c *Core) SignUpUser(ctx context.Context, email, password string) (user models.User, err error) {
	if _, err = mail.ParseAddress(email); err != nil {
		err = errs.ErrWrongEmail

		return
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return
	}

	return c.repo.AddUser(ctx, email, hashedPassword)
}

func (c *Core) SignInUser(ctx context.Context, email, password string) (token models.JWT, err error) {
	if _, err = mail.ParseAddress(email); err != nil {
		err = errs.ErrWrongEmail

		return
	}

	user, err := c.repo.GetUserByName(ctx, email, password)
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

	token.AccessTokenMaxAge = c.cfg.Security.AccessTokenMaxAge * minutesPerHour
	token.RefreshTokenMaxAge = c.cfg.Security.RefreshTokenMaxAge * minutesPerHour
	token.Domain = c.cfg.Security.Domain

	return token, nil
}

func (c *Core) RefreshAccessToken(ctx context.Context, refreshToken string) (token models.JWT, err error) {
	userID, err := helpers.ValidateToken(refreshToken, c.cfg.Security.RefreshTokenPublicKey)
	if err != nil {
		err = errs.ErrTokenValidation

		return
	}

	user, err := c.repo.GetUserByID(ctx, fmt.Sprint(userID))
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

	sub, err := helpers.ValidateToken(accessToken, c.cfg.AccessTokenPublicKey)
	if err != nil {
		err = errs.ErrTokenValidation

		return user, err
	}

	userID := fmt.Sprint(sub)
	user, err = c.repo.GetUserByID(ctx, userID)

	if err != nil {
		err = errs.ErrTokenValidation

		return user, err
	}

	return user, nil
}
