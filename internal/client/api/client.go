package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"github.com/vleukhin/GophKeeper/internal/models"
	"golang.org/x/exp/slices"
	"log"
	"net/http"
)

type Client interface {
	AuthClient
	CardsClient
	CredsClient
}

type AuthClient interface {
	Login(user *models.User) (models.JWT, error)
	Register(user *models.User) error
}

type CardsClient interface {
	GetCards(accessToken string) ([]models.Card, error)
	StoreCard(accessToken string, card *models.Card) error
	DelCard(accessToken, cardID string) error
}

type CredsClient interface {
	GetCreds(accessToken string) ([]models.Cred, error)
	AddCred(accessToken string, login *models.Cred) error
	DelCred(accessToken, loginID string) error
}

type HttpClient struct {
	host string
}

func NewHttpClient(host string) Client {
	return &HttpClient{host: host}
}

var errServer = errors.New("got server error")

func parseServerError(body []byte) string {
	var errMessage struct {
		Message string `json:"error"`
	}

	if err := json.Unmarshal(body, &errMessage); err == nil {
		return errMessage.Message
	}

	return ""
}

func (c *HttpClient) getEntities(entity interface{}, accessToken, endpoint string) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetResult(entity).
		Get(fmt.Sprintf("%s/%s", c.host, endpoint))
	if err != nil {
		log.Println(err)

		return err
	}

	if err := c.checkResCode(resp); err != nil {
		return err
	}

	return nil
}

func (c *HttpClient) delEntity(accessToken, endpoint, id string) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(fmt.Sprintf("%s/%s/%s", c.host, endpoint, id))
	if err != nil {
		log.Fatalf("GophKeeperClientAPI - client.R - %v ", err)
	}
	if err := c.checkResCode(resp); err != nil {
		return errServer
	}

	return nil
}

func (c *HttpClient) addEntity(entity interface{}, accessToken, endpoint string) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(entity).
		SetResult(entity).
		Post(fmt.Sprintf("%s/%s", c.host, endpoint))
	if err != nil {
		log.Fatalf("GophKeeperClientAPI - client.R - %v ", err)
	}
	if err := c.checkResCode(resp); err != nil {
		return errServer
	}

	return nil
}

func (c *HttpClient) checkResCode(resp *resty.Response) error {
	badCodes := []int{http.StatusBadRequest, http.StatusInternalServerError, http.StatusUnauthorized}
	if slices.Contains(badCodes, resp.StatusCode()) {
		errMessage := parseServerError(resp.Body())
		color.Red("Server error: %s", errMessage)

		return errServer
	}

	return nil
}
