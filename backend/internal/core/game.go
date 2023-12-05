package core

import (
	"context"
	"errors"
	"polygames/internal/domain"
	"polygames/internal/domain/errcore"
	"polygames/internal/repository/database"
)

func (c *core) CreateGame(ctx context.Context, req *domain.CreateGameRequest) (*domain.CreateGameResponse, error) {
	_, err := c.GetUser(ctx, &domain.GetUserRequest{UserId: req.UserID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	id, err := c.repo.CreateGame(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	response, err := c.GetGame(ctx, &domain.GetGameRequest{GameID: id})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.CreateGameResponse{Game: response.Game}, nil
}

func (c *core) GetGame(ctx context.Context, req *domain.GetGameRequest) (*domain.GetGameResponse, error) {
	game, err := c.repo.GetGame(ctx, req)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.GameNotFoundError
		}

		return nil, err
	}

	return &domain.GetGameResponse{Game: game}, nil
}

func (c *core) UpdateGame(ctx context.Context, req *domain.UpdateGameRequest) (*domain.UpdateGameResponse, error) {
	_, err := c.GetGame(ctx, &domain.GetGameRequest{GameID: req.ID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.UpdateGame(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	response, err := c.GetGame(ctx, &domain.GetGameRequest{GameID: req.ID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.UpdateGameResponse{Game: response.Game}, nil
}

func (c *core) DeleteGame(ctx context.Context, req *domain.DeleteGameRequest) (*domain.DeleteGameResponse, error) {
	_, err := c.GetGame(ctx, &domain.GetGameRequest{GameID: req.GameID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.DeleteGame(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.DeleteGameResponse{}, nil
}
