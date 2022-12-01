package helpers

import (
	"fmt"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func IsEmptyUUID(uuid uuid.UUID) bool {
	return uuid.ID() == 0
}
