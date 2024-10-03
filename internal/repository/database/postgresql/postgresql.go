package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"polygames/internal/repository/database"
)

type driver struct {
	conn *pgxpool.Pool
}

// NewDriver connects to PostgreSQL and return database.Database interface.
func NewDriver(ctx context.Context, connString string) (database.Database, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("connecting to postgresql: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("pinging postgresql: %w", err)
	}

	return &driver{conn: pool}, nil
}

func (d *driver) Close() {
	d.conn.Close()
}
