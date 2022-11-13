package storage

import "github.com/vleukhin/GophKeeper/internal/models"

type Repo interface {
	MigrateDB()

	UserRepo
}

type UserRepo interface {
	AddUser(user *models.User) error
	UpdateUserToken(user *models.User, token *models.JWT) error
	DropUserToken() error
	GetSavedAccessToken() (string, error)
	RemoveUsers()
	UserExistsByEmail(email string) bool
	GetUserPasswordHash() string
}
