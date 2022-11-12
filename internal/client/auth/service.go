package auth

import "fmt"

type Service struct {
}

func NewAuthService() *Service {
	return &Service{}
}

func (a Service) Login(login, password string) {
	fmt.Printf("attemt to log in with %s and %s", login, password)
}

func (a Service) Register(login, password string) {
	fmt.Printf("attemt to register with %s and %s", login, password)
}

func (a Service) Logout() {
	fmt.Printf("attemt to log out")
}
