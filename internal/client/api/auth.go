package api

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *HTTPClient) Login(user models.User) (models.JWT, error) {
	var token models.JWT
	payload := models.LoginPayload{
		Name:     user.Name,
		Password: user.Password,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return token, errors.Wrap(err, "failed to marshal user payload")
	}
	resp, err := c.post("api/auth/login", body, &token)
	if err != nil {
		return token, errors.Wrap(err, "request error")
	}
	if err := c.checkResCode(resp); err != nil {
		return token, err
	}

	return token, nil
}

func (c *HTTPClient) Register(user models.User) error {
	payload := models.LoginPayload{
		Name:     user.Name,
		Password: user.Password,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal user payload")
	}
	resp, err := c.post("api/auth/register", body, user)
	if err != nil {
		return errors.Wrap(err, "request error")
	}
	if err := c.checkResCode(resp); err != nil {
		return err
	}

	return nil
}
