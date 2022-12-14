package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/vleukhin/GophKeeper/internal/models"
)

type Client interface {
	AuthClient
	CardsClient
	CredsClient
	NoteClient
	FilesClient
}

type AuthClient interface {
	Login(name, password string) (models.User, models.JWT, error)
	Register(name, password string) (models.User, models.JWT, error)
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

type FilesClient interface {
	GetFiles(accessToken string) ([]models.File, error)
	StoreFile(accessToken string, file models.File, writer io.Reader) error
	DelFile(accessToken, fileID string) error
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
func (c *HTTPClient) delete(url string) (*resty.Response, error) {
	return c.client.R().
		SetHeader("Content-Type", "application/json").
		Delete(fmt.Sprintf("%s/%s", c.host, url))
}

func parseServerError(body []byte) string {
	var errResponse models.ErrMessage

	if err := json.Unmarshal(body, &errResponse); err == nil {
		return errResponse.Message
	}

	return "Unknown error"
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
	resp, err := c.delete(fmt.Sprintf("%s/%s", endpoint, id))
	if err != nil {
		return err
	}
	if err := c.checkResCode(resp); err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) addEntity(models interface{}, accessToken, endpoint string) error {
	c.client.SetAuthToken(accessToken)
	resp, err := c.post(endpoint, models, models)
	if err != nil {
		return err
	}
	if err := c.checkResCode(resp); err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) checkResCode(resp *resty.Response) error {
	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusAccepted {
		errMessage := parseServerError(resp.Body())
		return fmt.Errorf("server error: %s", errMessage)
	}

	return nil
}
