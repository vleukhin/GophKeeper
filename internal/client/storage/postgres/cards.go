package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/vleukhin/GophKeeper/internal/models"
)

const storeCardQuery = `
	INSERT INTO cards (id, name, card_holder_name, number, bank, exp_month, exp_year, security_code, meta)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

func (p Storage) StoreCard(ctx context.Context, card models.Card) error {
	_, err := p.conn.Exec(ctx, storeCardQuery, card.ID, card.Name, card.CardHolderName, card.Number, card.Bank,
		card.ExpirationMonth, card.ExpirationYear, card.SecurityCode, card.Meta)

	return err
}

func (p Storage) StoreCards(ctx context.Context, cards []models.Card) error {
	var err error
	for _, c := range cards {
		err = p.StoreCard(ctx, c)
		if err != nil {
			return err
		}
	}

	return nil
}

const getCardByIDQuery = `
	SELECT * FROM cards where id = $1
`

func (p Storage) GetCardByID(ctx context.Context, cardID uuid.UUID) (models.Card, error) {
	var result models.Card
	row := p.conn.QueryRow(ctx, getCardByIDQuery, cardID)
	err := row.Scan(&result.ID, &result.Name, &result.CardHolderName, &result.Number, &result.Bank,
		&result.ExpirationMonth, &result.ExpirationYear, &result.SecurityCode, &result.Meta)
	if err != nil && err != pgx.ErrNoRows {
		return result, err
	}
	return result, nil
}

const loadCardsQuery = `
	SELECT * from cards
`

func (p Storage) LoadCards(ctx context.Context) ([]models.Card, error) {
	var result []models.Card

	rows, err := p.conn.Query(ctx, loadCardsQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var card models.Card
		err := rows.Scan(&card.ID, &card.Name, &card.CardHolderName, &card.Number, &card.Bank,
			&card.ExpirationMonth, &card.ExpirationYear, &card.SecurityCode, &card.Meta)

		if err != nil {
			return nil, err
		}

		result = append(result, card)
	}

	return result, nil
}

const delCardQuery = `
	DELETE FROM cards WHERE id = $1
`

func (p Storage) DelCard(ctx context.Context, cardID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, delCardQuery, cardID)
	return err
}
