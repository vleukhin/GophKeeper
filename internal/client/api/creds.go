package api

import (
	"github.com/vleukhin/GophKeeper/internal/models"
)

const credsEndpoint = "api/creds"

func (c *HTTPClient) GetCreds(accessToken string) (logins []models.Cred, err error) {
	if err := c.getEntities(&logins, accessToken, credsEndpoint); err != nil {
		return nil, err
	}

	return logins, nil
}

func (c *HTTPClient) AddCred(accessToken string, login *models.Cred) error {
	return c.addEntity(login, accessToken, credsEndpoint)
}

func (c *HTTPClient) DelCred(accessToken, loginID string) error {
	return c.delEntity(accessToken, credsEndpoint, loginID)
}
