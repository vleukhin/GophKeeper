package api

import (
	"github.com/vleukhin/GophKeeper/internal/models"
)

const loginsEndpoint = "api/v1/user/logins"

func (c *HttpClient) GetCreds(accessToken string) (logins []models.Cred, err error) {
	if err := c.getEntities(&logins, accessToken, loginsEndpoint); err != nil {
		return nil, err
	}

	return logins, nil
}

func (c *HttpClient) AddCred(accessToken string, login *models.Cred) error {
	return c.addEntity(login, accessToken, loginsEndpoint)
}

func (c *HttpClient) DelCred(accessToken, loginID string) error {
	return c.delEntity(accessToken, loginsEndpoint, loginID)
}
