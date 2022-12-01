package core

import (
	"context"
	"fmt"
	"log"

	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func (c *Core) StoreCard(card *models.Card) {
	accessToken, err := c.authorisationCheck()
	if err != nil {
		return
	}

	c.encryptCard(c.cfg.EncryptKey, card)

	if err = c.client.StoreCard(accessToken, card); err != nil {
		return
	}

	if err = c.storage.StoreCard(context.Background(), *card); err != nil {
		log.Fatal(err)
	}

	color.Green("Card %q added, id: %v", card.Name, card.ID)
}

func (c *Core) encryptCard(key string, card *models.Card) {
	card.Number = helpers.Encrypt(key, card.Number)
	card.SecurityCode = helpers.Encrypt(key, card.SecurityCode)
	card.ExpirationMonth = helpers.Encrypt(key, card.ExpirationMonth)
	card.ExpirationYear = helpers.Encrypt(key, card.ExpirationYear)
	card.CardHolderName = helpers.Encrypt(key, card.CardHolderName)
}

func (c *Core) decryptCard(key string, card *models.Card) {
	card.Number = helpers.Decrypt(key, card.Number)
	card.SecurityCode = helpers.Decrypt(key, card.SecurityCode)
	card.ExpirationMonth = helpers.Decrypt(key, card.ExpirationMonth)
	card.ExpirationYear = helpers.Decrypt(key, card.ExpirationYear)
	card.CardHolderName = helpers.Decrypt(key, card.CardHolderName)
}

func (c *Core) loadCards(accessToken string) {
	cards, err := c.client.GetCards(accessToken)
	if err != nil {
		color.Red("Connection error: %v", err)

		return
	}

	if err = c.storage.StoreCards(context.Background(), cards); err != nil {
		log.Println(err)

		return
	}
	color.Green("Loaded %v cards", len(cards))
}

func (c *Core) ShowCard(cardID string) {
	cardUUID, err := uuid.Parse(cardID)
	if err != nil {
		color.Red(err.Error())

		return
	}
	card, err := c.storage.GetCardByID(context.Background(), cardUUID)
	if err != nil {
		color.Red(err.Error())

		return
	}
	c.decryptCard(c.cfg.EncryptKey, &card)
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("ID: %s\nname:%s\nCardHolderName:%s\nNumber:%s\nBrand:%s\nExpiration: %s/%s\nCode%s\n%v\n", //nolint:forbidigo // cli printing
		yellow(card.ID),
		yellow(card.Name),
		yellow(card.CardHolderName),
		yellow(card.Number),
		yellow(card.Bank),
		yellow(card.ExpirationMonth),
		yellow(card.ExpirationYear),
		yellow(card.SecurityCode),
		yellow(card.Meta),
	)
}

func (c *Core) DelCard(cardID string) {
	accessToken, err := c.authorisationCheck()
	if err != nil {
		return
	}
	cardUUID, err := uuid.Parse(cardID)
	if err != nil {
		color.Red(err.Error())
		log.Fatalf("Core - uuid.Parse - %v", err)
	}

	if err := c.storage.DelCard(context.Background(), cardUUID); err != nil {
		log.Fatalf("Core - storage.DelCard - %v", err)
	}

	if err := c.client.DelCard(accessToken, cardID); err != nil {
		log.Fatalf("Core - storage.DelCard - %v", err)
	}

	color.Green("Card %q removed", cardID)
}
