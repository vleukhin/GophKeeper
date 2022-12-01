package api

import (
	"encoding/json"
	"github.com/pkg/errors"

	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *HTTPClient) Login(name, password string) (models.JWT, error) {
	var token models.JWT
	payload := models.LoginPayload{
		Name:     name,
		Password: password,
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

func (c *HTTPClient) Register(name, password string) (models.User, models.JWT, error) {
	var response models.RegisterResponse
	payload := models.LoginPayload{
		Name:     name,
		Password: password,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return models.User{}, models.JWT{}, err
	}
	resp, err := c.post("api/auth/register", body, &response)
	if err != nil {
		return models.User{}, models.JWT{}, err
	}
	if err := c.checkResCode(resp); err != nil {
		return models.User{}, models.JWT{}, err
	}

	return response.User, response.Token, nil
}
