package api

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"github.com/vleukhin/GophKeeper/internal/models"
	"log"
	"net/http"
)

func (c *HttpClient) Login(user *models.User) (token models.JWT, err error) {
	client := resty.New()
	body := fmt.Sprintf(`{"email":%q, "password":%q}`, user.Email, user.Password)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&token).
		Post(fmt.Sprintf("%s/api/login", c.host))
	if err != nil {
		return
	}

	if resp.StatusCode() == http.StatusBadRequest || resp.StatusCode() == http.StatusInternalServerError {
		color.Red("Server error: %s", parseServerError(resp.Body()))

		return token, errServer
	}

	return token, nil
}

func (c *HttpClient) Register(user *models.User) error {
	client := resty.New()
	body := fmt.Sprintf(`{"email":%q, "password":%q}`, user.Email, user.Password)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(user).
		Post(fmt.Sprintf("%s/api/register", c.host))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() == http.StatusBadRequest || resp.StatusCode() == http.StatusInternalServerError {
		errMessage := parseServerError(resp.Body())
		color.Red("Server error: %s", errMessage)

		return errServer
	}

	return nil
}
