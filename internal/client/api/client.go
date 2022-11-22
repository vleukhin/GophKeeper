package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slices"

	"github.com/vleukhin/GophKeeper/internal/models"
)

type Client interface {
	AuthClient
	CardsClient
	CredsClient
	NoteClient
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

type NoteClient interface {
	GetNotes(accessToken string) ([]models.Note, error)
	StoreNote(accessToken string, note *models.Note) error
	DelNote(accessToken, noteID string) error
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

func (c *HttpClient) getEntities(models interface{}, accessToken, endpoint string) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetResult(models).
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

func (c *HttpClient) addEntity(models interface{}, accessToken, endpoint string) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(models).
		SetResult(models).
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
