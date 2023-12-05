package database

import (
	"context"
	"polygames/internal/domain"
)

type Database interface {
	CreateGame(ctx context.Context, req *domain.CreateGameRequest) (int64, error)
	GetGame(ctx context.Context, req *domain.GetGameRequest) (*domain.Game, error)
	UpdateGame(ctx context.Context, req *domain.UpdateGameRequest) error
	DeleteGame(ctx context.Context, req *domain.DeleteGameRequest) error

	CreateGenre(ctx context.Context, req *domain.CreateGenreRequest) (int64, error)
	GetGenre(ctx context.Context, req *domain.GetGenreRequest) (*domain.Genre, error)
	UpdateGenre(ctx context.Context, req *domain.UpdateGenreRequest) error
	DeleteGenre(ctx context.Context, req *domain.DeleteGenreRequest) error

	GetUser(ctx context.Context, req *domain.GetUserRequest) (*domain.User, error)
	GetUserByLogin(ctx context.Context, login string) (*domain.User, error)
	CheckUsernameUniqueness(ctx context.Context, username string) error
	CheckEmailUniqueness(ctx context.Context, email string) error
	DeleteUser(ctx context.Context, req *domain.DeleteUserRequest) error

	SignUp(ctx context.Context, req *domain.SignUpRequest) (int64, error)

	Close()
}
