package auth

type UserAuth interface {
	Login(login, password string)
}

type Service struct {
}

func NewAuthService() *Service {
	return &Service{}
}
