package core

import (
	"context"
	"errors"
	"polygames/internal/domain"
	"polygames/internal/domain/errcore"
	"polygames/internal/repository/database"
)

func (c *core) CreateGenre(ctx context.Context, req *domain.CreateGenreRequest) (*domain.CreateGenreResponse, error) {
	id, err := c.repo.CreateGenre(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	response, err := c.GetGenre(ctx, &domain.GetGenreRequest{GenreID: id})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.CreateGenreResponse{Genre: response.Genre}, nil
}

func (c *core) GetGenre(ctx context.Context, req *domain.GetGenreRequest) (*domain.GetGenreResponse, error) {
	genre, err := c.repo.GetGenre(ctx, req)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.GenreNotFoundError
		}

		return nil, err
	}

	return &domain.GetGenreResponse{Genre: genre}, nil
}

func (c *core) UpdateGenre(ctx context.Context, req *domain.UpdateGenreRequest) (*domain.UpdateGenreResponse, error) {
	_, err := c.GetGenre(ctx, &domain.GetGenreRequest{GenreID: req.ID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.UpdateGenre(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	response, err := c.GetGenre(ctx, &domain.GetGenreRequest{GenreID: req.ID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.UpdateGenreResponse{Genre: response.Genre}, nil
}

func (c *core) DeleteGenre(ctx context.Context, req *domain.DeleteGenreRequest) (*domain.DeleteGenreResponse, error) {
	_, err := c.GetGenre(ctx, &domain.GetGenreRequest{GenreID: req.GenreID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.DeleteGenre(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.DeleteGenreResponse{}, nil
}
