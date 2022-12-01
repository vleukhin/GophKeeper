package postgres

import (
	"context"
	"fmt"
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

func (p Storage) Migrate(ctx context.Context) error {
	migrations := []string{
		createUsersTableQuery,
		createCredsTableQuery,
		createCardsTableQuery,
		createNotesTableQuery,
		createFilesTableQuery,
	}
	for _, m := range migrations {
		fmt.Println(m)
		_, err := p.conn.Query(ctx, m)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p Storage) Ping(ctx context.Context) error {
	return p.conn.Ping(ctx)
}
