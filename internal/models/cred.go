package models

import "github.com/google/uuid"

type Cred struct {
	ID       uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	URI      string    `json:"uri"`
	Meta     []Meta    `json:"meta"`
}
