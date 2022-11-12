package storage

import entity "github.com/vleukhin/GophKeeper/internal/models"

type Client interface {
	GetCards(accessToken string) ([]entity.Card, error)
	StoreCard(accessToken string, card *entity.Card) error
	DelCard(accessToken, cardID string) error
}
