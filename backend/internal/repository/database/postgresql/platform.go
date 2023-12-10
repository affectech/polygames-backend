package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"polygames/internal/domain"
	"polygames/internal/repository/database"
)

func (d *driver) CreatePlatform(ctx context.Context, req *domain.CreatePlatformRequest) (int64, error) {
	row := d.conn.QueryRow(ctx, `insert into platforms(name)
                                     values($1)
                                     returning id`,
	)
	var PlatformID int64

	if err := row.Scan(&PlatformID); err != nil {
		return 0, fmt.Errorf("%w: scanning platform_id", err)
	}

	return PlatformID, nil
}

func (d *driver) GetPlatform(ctx context.Context, req *domain.GetPlatformRequest) (*domain.Platform, error) {
	row := d.conn.QueryRow(ctx, `select 
                                     c.id
                                     from platforms c
                                     where c.id = $1`, req.PlatformID)

	var Platform domain.Platform
	if err := row.Scan(
		&Platform.ID,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrObjectNotFound
		}
		return nil, fmt.Errorf("%w: scanning platform", err)
	}

	return &Platform, nil
}

func (d *driver) UpdatePlatform(ctx context.Context, req *domain.UpdatePlatformRequest) error {
	var ok bool
	err := d.conn.QueryRow(ctx, `update platforms 
                                     set name = $2
                                     where id = $1
                                     returning true`,
		req.ID,
		req.Name,
	).Scan(&ok)
	if err != nil {
		return fmt.Errorf("updating platform: %w", err)
	}

	return nil
}

func (d *driver) DeletePlatform(ctx context.Context, req *domain.DeletePlatformRequest) error {
	var ok bool
	err := d.conn.QueryRow(ctx, `delete from platforms 
                                     where id = $1
                                     returning true`,
		req.PlatformID,
	).Scan(&ok)
	if err != nil {
		return fmt.Errorf("deleting platform: %w", err)
	}

	return nil
}
