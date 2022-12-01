package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Password string    `json:"-"`
}

type JWT struct {
	AccessToken        string        `json:"access_token"`
	RefreshToken       string        `json:"refresh_token"`
	AccessTokenMaxAge  time.Duration `json:"-"`
	RefreshTokenMaxAge time.Duration `json:"-"`
	Domain             string        `json:"-"`
}

type Card struct {
	ID              uuid.UUID `json:"uuid"`
	Name            string    `json:"name"`
	CardHolderName  string    `json:"card_holder_name"`
	Number          string    `json:"number"`
	Bank            string    `json:"bank"`
	ExpirationMonth string    `json:"expiration_month"`
	ExpirationYear  string    `json:"expiration_year"`
	SecurityCode    string    `json:"security_code"`
	Meta            []Meta    `json:"meta"`
}

type Cred struct {
	ID       uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	URI      string    `json:"uri"`
	Meta     []Meta    `json:"meta"`
}

type Note struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
	Text string    `json:"text"`
	Meta []Meta    `json:"meta"`
}

type File struct {
	ID       uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	FileName string    `json:"file_name"`
	Content  []byte    `json:"content"`
	Meta     []Meta    `json:"meta"`
}

type Meta struct {
	ID    uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Value string    `json:"value"`
}
