package models

type LoginPayload struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ErrMessage struct {
	Message string `json:"error"`
}
