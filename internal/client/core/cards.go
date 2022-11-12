package core

import (
	"fmt"
	"github.com/vleukhin/GophKeeper/internal/helpers"
	"github.com/vleukhin/GophKeeper/internal/models"
	"log"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func (c *Core) StoreCard(userPassword string, card *models.Card) {
	accessToken, err := c.authorisationCheck(userPassword)
	if err != nil {
		return
	}

	c.encryptCard(userPassword, card)

	if err = c.client.StoreCard(accessToken, card); err != nil {
		return
	}

	if err = c.repo.StoreCard(card); err != nil {
		log.Fatal(err)
	}

	color.Green("Card %q added, id: %v", card.Name, card.ID)
}

func (c *Core) encryptCard(userPassword string, card *models.Card) {
	card.Number = helpers.Encrypt(userPassword, card.Number)
	card.SecurityCode = helpers.Encrypt(userPassword, card.SecurityCode)
	card.ExpirationMonth = helpers.Encrypt(userPassword, card.ExpirationMonth)
	card.ExpirationYear = helpers.Encrypt(userPassword, card.ExpirationYear)
	card.CardHolderName = helpers.Encrypt(userPassword, card.CardHolderName)
}

func (c *Core) decryptCard(userPassword string, card *models.Card) {
	card.Number = helpers.Decrypt(userPassword, card.Number)
	card.SecurityCode = helpers.Decrypt(userPassword, card.SecurityCode)
	card.ExpirationMonth = helpers.Decrypt(userPassword, card.ExpirationMonth)
	card.ExpirationYear = helpers.Decrypt(userPassword, card.ExpirationYear)
	card.CardHolderName = helpers.Decrypt(userPassword, card.CardHolderName)
}

func (c *Core) loadCards(accessToken string) {
	cards, err := c.client.GetCards(accessToken)
	if err != nil {
		color.Red("Connection error: %v", err)

		return
	}

	if err = c.repo.StoreCards(cards); err != nil {
		log.Println(err)

		return
	}
	color.Green("Loaded %v cards", len(cards))
}

func (c *Core) ShowCard(userPassword, cardID string) {
	if !c.verifyPassword(userPassword) {
		return
	}
	cardUUID, err := uuid.Parse(cardID)
	if err != nil {
		color.Red(err.Error())

		return
	}
	card, err := c.repo.GetCardByID(cardUUID)
	if err != nil {
		color.Red(err.Error())

		return
	}
	c.decryptCard(userPassword, &card)
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("ID: %s\nname:%s\nCardHolderName:%s\nNumber:%s\nBrand:%s\nExpiration: %s/%s\nCode%s\n%v\n", //nolint:forbidigo // cli printing
		yellow(card.ID),
		yellow(card.Name),
		yellow(card.CardHolderName),
		yellow(card.Number),
		yellow(card.Brand),
		yellow(card.ExpirationMonth),
		yellow(card.ExpirationYear),
		yellow(card.SecurityCode),
		yellow(card.Meta),
	)
}

func (c *Core) DelCard(userPassword, cardID string) {
	accessToken, err := c.authorisationCheck(userPassword)
	if err != nil {
		return
	}
	cardUUID, err := uuid.Parse(cardID)
	if err != nil {
		color.Red(err.Error())
		log.Fatalf("GophKeeperClientUseCase - uuid.Parse - %v", err)
	}

	if err := c.repo.DelCard(cardUUID); err != nil {
		log.Fatalf("GophKeeperClientUseCase - repo.DelCard - %v", err)
	}

	if err := c.client.DelCard(accessToken, cardID); err != nil {
		log.Fatalf("GophKeeperClientUseCase - repo.DelCard - %v", err)
	}

	color.Green("Card %q removed", cardID)
}
