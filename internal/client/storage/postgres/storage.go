package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	conn *pgxpool.Pool
}

func NewPostgresStorage(dsn string, connTimeout time.Duration) (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	conn, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Storage{
		conn: conn,
	}, nil
}

func (p Storage) MigrateDB(ctx context.Context) error {
	migrations := []string{
		createUsersTableQuery,
		createCredsTableQuery,
		createCardsTableQuery,
		createNotesTableQuery,
	}
	for _, m := range migrations {
		_, err := p.conn.Query(ctx, m)
		if err != nil {
			return err
		}
	}

	return nil
}
