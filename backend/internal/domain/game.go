package domain

import (
	"time"
)

type (
	Game struct {
		ID          int64      `json:"id"`
		UserID      int64      `json:"user_id"`
		TeamID      int64      `json:"team_id"`
		Description string     `json:"description"`
		GenreID     int64      `json:"genre_id"`
		ImageID     string     `json:"image_id"`
		FileID      string     `json:"file_id"`
		Likes       int64      `json:"likes"`
		Views       int64      `json:"views"`
		Downloads   int64      `json:"downloads"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	}
	CreateGameRequest struct {
		ID          int64  `json:"id"`
		UserID      int64  `json:"user_id"`
		TeamID      int64  `json:"team_id"`
		Description string `json:"description"`
		GenreID     int64  `json:"genre_id"`
		ImageID     string `json:"image_id"`
		FileID      string `json:"file_id"`
	}

	CreateGameResponse struct {
		Game *Game `json:"data"`
	}
	GetGameRequest struct {
		GameID int64 `json:"game_id"`
	}
	GetGameResponse struct {
		Game *Game `json:"data"`
	}
	UpdateGameRequest struct {
		ID          int64  `json:"id"`
		UserID      int64  `json:"user_id"`
		TeamID      int64  `json:"team_id"`
		Description string `json:"description"`
		GenreID     int64  `json:"genre_id"`
		ImageID     string `json:"image_id"`
		FileID      string `json:"file_id"`
	}
	UpdateGameResponse struct {
		Game *Game `json:"data"`
	}
	DeleteGameRequest struct {
		GameID int64 `json:"game_id"`
	}
	DeleteGameResponse struct {
	}
)
