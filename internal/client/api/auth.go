package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"

	"github.com/vleukhin/GophKeeper/internal/models"
)

func (c *HTTPClient) Login(user *models.User) (token models.JWT, err error) {
	client := resty.New()
	body := fmt.Sprintf(`{"name":%q, "password":%q}`, user.Name, user.Password)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&token).
		Post(fmt.Sprintf("%s/api/auth/login", c.host))
	if err != nil {
		return
	}

	if resp.StatusCode() == http.StatusBadRequest || resp.StatusCode() == http.StatusInternalServerError {
		color.Red("Server error: %s", parseServerError(resp.Body()))

		return token, errServer
	}

	return token, nil
}

func (c *HTTPClient) Register(user *models.User) error {
	client := resty.New()
	body := fmt.Sprintf(`{"name":%q, "password":%q}`, user.Name, user.Password)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(user).
		Post(fmt.Sprintf("%s/api/auth/register", c.host))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		errMessage := parseServerError(resp.Body())
		color.Red("Server error: %s \nStatus code: %d", errMessage, resp.StatusCode())

		return errServer
	}

	return nil
}
