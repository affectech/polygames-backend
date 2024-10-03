package core

import (
	"context"
	"fmt"
	"polygames/internal/domain"
	"polygames/internal/infrastructure/config"
	"polygames/internal/repository/database"
	"polygames/internal/repository/database/postgresql"
)

// Core represents business logic layer interface.
type Core interface {
	CreateGame(ctx context.Context, req *domain.CreateGameRequest) (*domain.CreateGameResponse, error)
	GetGame(ctx context.Context, req *domain.GetGameRequest) (*domain.GetGameResponse, error)
	UpdateGame(ctx context.Context, req *domain.UpdateGameRequest) (*domain.UpdateGameResponse, error)
	DeleteGame(ctx context.Context, req *domain.DeleteGameRequest) (*domain.DeleteGameResponse, error)

	CreateGenre(ctx context.Context, req *domain.CreateGenreRequest) (*domain.CreateGenreResponse, error)
	GetGenre(ctx context.Context, req *domain.GetGenreRequest) (*domain.GetGenreResponse, error)
	UpdateGenre(ctx context.Context, req *domain.UpdateGenreRequest) (*domain.UpdateGenreResponse, error)
	DeleteGenre(ctx context.Context, req *domain.DeleteGenreRequest) (*domain.DeleteGenreResponse, error)

	CreatePlatform(ctx context.Context, req *domain.CreatePlatformRequest) (*domain.CreatePlatformResponse, error)
	GetPlatform(ctx context.Context, req *domain.GetPlatformRequest) (*domain.GetPlatformResponse, error)
	UpdatePlatform(ctx context.Context, req *domain.UpdatePlatformRequest) (*domain.UpdatePlatformResponse, error)
	DeletePlatform(ctx context.Context, req *domain.DeletePlatformRequest) (*domain.DeletePlatformResponse, error)

	SignIn(ctx context.Context, req *domain.SignInRequest) (*domain.SignInResponse, error)
	SignUp(ctx context.Context, req *domain.SignUpRequest) (*domain.SignUpResponse, error)

	GetUser(ctx context.Context, req *domain.GetUserRequest) (*domain.GetUserResponse, error)
	DeleteUser(ctx context.Context, req *domain.DeleteUserRequest) (*domain.DeleteUserResponse, error)
}

// core implements Core interface.
type core struct {
	repo database.Database
}

// New returns Core instance.
func New(ctx context.Context) (Core, error) {
	db, err := postgresql.NewDriver(ctx, config.Config.DatabaseConnString)
	if err != nil {
		return nil, fmt.Errorf("creating postgresql driver: %w", err)
	}

	return &core{repo: db}, nil
}
