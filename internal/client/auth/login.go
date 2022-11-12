package auth

import "fmt"

func (a Service) Login(login, password string) {
	fmt.Printf("attemt to log in with %s and %s", login, password)
}
