package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"github.com/vleukhin/GophKeeper/internal/models"
	"log"
	"net/http"
)

type Client interface {
	Login(user *models.User) (models.JWT, error)
	Register(user *models.User) error
}

type HttpClient struct {
	host string
}

func NewHttpClient(host string) *HttpClient {
	return &HttpClient{host: host}
}

var errServer = errors.New("got server error")

func (c *HttpClient) Login(user *models.User) (token models.JWT, err error) {
	client := resty.New()
	body := fmt.Sprintf(`{"email":%q, "password":%q}`, user.Email, user.Password)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&token).
		Post(fmt.Sprintf("%s/api/v1/auth/login", c.host))
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
		Post(fmt.Sprintf("%s/api/v1/auth/register", c.host))
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

func parseServerError(body []byte) string {
	var errMessage struct {
		Message string `json:"error"`
	}

	if err := json.Unmarshal(body, &errMessage); err == nil {
		return errMessage.Message
	}

	return ""
}
