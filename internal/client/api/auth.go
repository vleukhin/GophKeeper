package api

import (
	"encoding/json"

	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *HTTPClient) Login(name, password string) (models.User, models.JWT, error) {
	var response models.AuthResponse
	payload := models.LoginPayload{
		Name:     name,
		Password: password,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return models.User{}, models.JWT{}, err
	}
	resp, err := c.post("api/auth/login", body, &response)
	if err != nil {
		return models.User{}, models.JWT{}, err
	}
	if err := c.checkResCode(resp); err != nil {
		return models.User{}, models.JWT{}, err
	}

	return response.User, response.Token, nil
}

func (c *HTTPClient) Register(name, password string) (models.User, models.JWT, error) {
	var response models.AuthResponse
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
