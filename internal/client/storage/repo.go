package storage

type Repo interface {
	MigrateDB()

	GetUserPasswordHash() string
	GetSavedAccessToken() (string, error)
	UserExistsByEmail(email string) bool
}
