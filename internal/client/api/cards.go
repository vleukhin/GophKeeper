package api

import "github.com/vleukhin/GophKeeper/internal/models"

const cardsEndpoint = "cards"

func (c *HttpClient) GetCards(accessToken string) (cards []models.Card, err error) {
	if err := c.getEntities(&cards, accessToken, cardsEndpoint); err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *HttpClient) StoreCard(accessToken string, card *models.Card) error {
	return c.addEntity(card, accessToken, cardsEndpoint)
}

func (c *HttpClient) DelCard(accessToken, cardID string) error {
	return c.delEntity(accessToken, cardsEndpoint, cardID)
}
