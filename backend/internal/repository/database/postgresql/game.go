package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"polygames/internal/domain"
	"polygames/internal/repository/database"
)

func (d *driver) CreateGame(ctx context.Context, req *domain.CreateGameRequest) (int64, error) {
	row := d.conn.QueryRow(ctx, `insert into games(user_id, team_id, description, genre_id, image_id, file_id)
                                     values($1, $2, $3, $4, $5, $6)
                                     returning id`,
		req.UserID,
		req.TeamID,
		req.Description,
		req.GenreID,
		req.ImageID,
		req.FileID,
	)
	var GameID int64

	if err := row.Scan(&GameID); err != nil {
		return 0, fmt.Errorf("%w: scanning game_id", err)
	}

	return GameID, nil
}

func (d *driver) GetGame(ctx context.Context, req *domain.GetGameRequest) (*domain.Game, error) {
	row := d.conn.QueryRow(ctx, `select 
                                     c.id, c.user_id, c.team_id, c.description, c.genre_id, c.image_id, c.file_id, c.created_at, c.updated_at
                                     from games c
                                     where c.id = $1`, req.GameID)

	var Game domain.Game
	if err := row.Scan(
		&Game.ID,
		&Game.UserID,
		&Game.TeamID,
		&Game.Description,
		&Game.GenreID,
		&Game.ImageID,
		&Game.FileID,
		&Game.CreatedAt,
		&Game.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrObjectNotFound
		}
		return nil, fmt.Errorf("%w: scanning game", err)
	}

	return &Game, nil
}

func (d *driver) UpdateGame(ctx context.Context, req *domain.UpdateGameRequest) error {
	var ok bool
	err := d.conn.QueryRow(ctx, `update games 
                                     set description = $4, genre_id = $5, image_id = $6, file_id = $7, updated_at = now()
                                     where id = $1
                                     returning true`,
		req.ID,
		req.UserID,
		req.TeamID,
		req.Description,
		req.GenreID,
		req.ImageID,
		req.FileID,
	).Scan(&ok)
	if err != nil {
		return fmt.Errorf("updating game: %w", err)
	}

	return nil
}

func (d *driver) DeleteGame(ctx context.Context, req *domain.DeleteGameRequest) error {
	var ok bool
	err := d.conn.QueryRow(ctx, `delete from games 
                                     where id = $1
                                     returning true`,
		req.GameID,
	).Scan(&ok)
	if err != nil {
		return fmt.Errorf("deleting game: %w", err)
	}

	return nil
}
