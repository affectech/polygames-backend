package core

import (
	"context"
	"errors"
	"polygames/internal/domain"
	"polygames/internal/domain/errcore"
	"polygames/internal/repository/database"
)

func (c *core) CreatePlatform(ctx context.Context, req *domain.CreatePlatformRequest) (*domain.CreatePlatformResponse, error) {
	id, err := c.repo.CreatePlatform(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	response, err := c.GetPlatform(ctx, &domain.GetPlatformRequest{PlatformID: id})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.CreatePlatformResponse{Platform: response.Platform}, nil
}

func (c *core) GetPlatform(ctx context.Context, req *domain.GetPlatformRequest) (*domain.GetPlatformResponse, error) {
	Platform, err := c.repo.GetPlatform(ctx, req)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.PlatformNotFoundError
		}

		return nil, err
	}

	return &domain.GetPlatformResponse{Platform: Platform}, nil
}

func (c *core) UpdatePlatform(ctx context.Context, req *domain.UpdatePlatformRequest) (*domain.UpdatePlatformResponse, error) {
	_, err := c.GetPlatform(ctx, &domain.GetPlatformRequest{PlatformID: req.ID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.UpdatePlatform(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	response, err := c.GetPlatform(ctx, &domain.GetPlatformRequest{PlatformID: req.ID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.UpdatePlatformResponse{Platform: response.Platform}, nil
}

func (c *core) DeletePlatform(ctx context.Context, req *domain.DeletePlatformRequest) (*domain.DeletePlatformResponse, error) {
	_, err := c.GetPlatform(ctx, &domain.GetPlatformRequest{PlatformID: req.PlatformID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.DeletePlatform(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.DeletePlatformResponse{}, nil
}
