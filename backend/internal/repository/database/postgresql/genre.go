package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"polygames/internal/domain"
	"polygames/internal/repository/database"
)

func (d *driver) CreateGenre(ctx context.Context, req *domain.CreateGenreRequest) (int64, error) {
	row := d.conn.QueryRow(ctx, `insert into genres(genre_id)
                                     values($1)
                                     returning id`,
	)
	var GenreID int64

	if err := row.Scan(&GenreID); err != nil {
		return 0, fmt.Errorf("%w: scanning genre_id", err)
	}

	return GenreID, nil
}

func (d *driver) GetGenre(ctx context.Context, req *domain.GetGenreRequest) (*domain.Genre, error) {
	row := d.conn.QueryRow(ctx, `select 
                                     c.id
                                     from genres c
                                     where c.id = $1`, req.GenreID)

	var Genre domain.Genre
	if err := row.Scan(
		&Genre.ID,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrObjectNotFound
		}
		return nil, fmt.Errorf("%w: scanning genre", err)
	}

	return &Genre, nil
}

func (d *driver) UpdateGenre(ctx context.Context, req *domain.UpdateGenreRequest) error {
	var ok bool
	err := d.conn.QueryRow(ctx, `update genres 
                                     set name = $2
                                     where id = $1
                                     returning true`,
		req.ID,
		req.Name,
	).Scan(&ok)
	if err != nil {
		return fmt.Errorf("updating genre: %w", err)
	}

	return nil
}

func (d *driver) DeleteGenre(ctx context.Context, req *domain.DeleteGenreRequest) error {
	var ok bool
	err := d.conn.QueryRow(ctx, `delete from genres 
                                     where id = $1
                                     returning true`,
		req.GenreID,
	).Scan(&ok)
	if err != nil {
		return fmt.Errorf("deleting genre: %w", err)
	}

	return nil
}
