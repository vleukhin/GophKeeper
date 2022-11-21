package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
	"github.com/vleukhin/GophKeeper/internal/models"
)

const getCardsQuery = `
	SELECT id, name, card_holder_name, number, bank, exp_month, exp_year, security_code, meta
	FROM cards where user_id = $1
`

func (p Storage) GetCards(ctx context.Context, user models.User) ([]models.Card, error) {
	var result []models.Card

	rows, err := p.conn.Query(ctx, getCardsQuery, user.ID)
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

const storeCardQuery = `
	INSERT INTO cards (id, user_id, name, card_holder_name, number, bank, exp_month, exp_year, security_code, meta)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

func (p Storage) AddCard(ctx context.Context, card *models.Card, userID uuid.UUID) error {
	_, err := p.conn.Exec(ctx, storeCardQuery, card.ID, userID, card.Name, card.CardHolderName, card.Number, card.Bank,
		card.ExpirationMonth, card.ExpirationYear, card.SecurityCode, card.Meta)

	return err
}

const delCardQuery = `
	DELETE FROM cards WHERE id = $1
`

func (p Storage) DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error {
	if !p.IsCardOwner(ctx, cardUUID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}
	_, err := p.conn.Exec(ctx, delCardQuery, cardUUID)
	return err
}

const updateCardQuery = `
	UPDATE cards SET
		 name = $1, 
		 card_holder_name = $2,
		 number = $3, 
		 bank = $4,
		 exp_month = $5,
		 exp_year = $6,
		 security_code = $7,
		 meta = $8
	WHERE id = $9
`

func (p Storage) UpdateCard(ctx context.Context, card *models.Card, userID uuid.UUID) error {
	if !p.IsCardOwner(ctx, card.ID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}
	_, err := p.conn.Exec(ctx, updateCardQuery,
		card.Name,
		card.CardHolderName,
		card.Number,
		card.Bank,
		card.ExpirationYear,
		card.ExpirationYear,
		card.SecurityCode,
		card.Meta,
		card.ID,
	)

	return err
}

const getCardByID = `
	SELECT id FROM cards WHERE id = $1 and user_id = $2
`

func (p Storage) IsCardOwner(ctx context.Context, cardUUID, userID uuid.UUID) bool {
	return p.conn.QueryRow(ctx, getCardByID, cardUUID, userID).Scan() == nil
}
