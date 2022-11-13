package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"uuid"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

type JWT struct {
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	AccessTokenMaxAge  int    `json:"-"`
	RefreshTokenMaxAge int    `json:"-"`
	Domain             string `json:"-"`
}
