package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
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
	Login(user models.User) (models.JWT, error)
	Register(user models.User) error
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

type HTTPClient struct {
	host   string
	client *resty.Client
}

func NewHTTPClient(host string) Client {
	return &HTTPClient{
		host:   host,
		client: resty.New(),
	}
}

var errServer = errors.New("got server error")

func (c *HTTPClient) post(url string, body, result interface{}) (*resty.Response, error) {
	return c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&result).
		Post(fmt.Sprintf("%s/%s", c.host, url))
}
func (c *HTTPClient) get(url string, result interface{}) (*resty.Response, error) {
	return c.client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&result).
		Get(fmt.Sprintf("%s/%s", c.host, url))
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

func (c *HTTPClient) getEntities(models interface{}, accessToken, endpoint string) error {
	c.client.SetAuthToken(accessToken)
	resp, err := c.get(endpoint, models)
	if err != nil {
		return errors.Wrap(err, "request error")
	}
	if err := c.checkResCode(resp); err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) delEntity(accessToken, endpoint, id string) error {
	c.client.SetAuthToken(accessToken)
	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		Delete(fmt.Sprintf("%s/%s/%s", c.host, endpoint, id))
	if err != nil {
		return err
	}
	if err := c.checkResCode(resp); err != nil {
		return errServer
	}

	return nil
}

func (c *HTTPClient) addEntity(models interface{}, accessToken, endpoint string) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(models).
		SetResult(models).
		Post(fmt.Sprintf("%s/%s", c.host, endpoint))
	if err != nil {
		return err
	}
	if err := c.checkResCode(resp); err != nil {
		return errServer
	}

	return nil
}

func (c *HTTPClient) checkResCode(resp *resty.Response) error {
	badCodes := []int{http.StatusBadRequest, http.StatusInternalServerError, http.StatusUnauthorized}
	if slices.Contains(badCodes, resp.StatusCode()) {
		errMessage := parseServerError(resp.Body())
		return errors.New(fmt.Sprintf("Server error: %s", errMessage))
	}

	return nil
}
