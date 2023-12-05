package core

import (
	"context"
	"errors"
	"polygames/internal/domain"
	"polygames/internal/domain/errcore"
	"polygames/internal/repository/database"
)

func (c *core) GetUser(ctx context.Context, req *domain.GetUserRequest) (*domain.GetUserResponse, error) {
	user, err := c.repo.GetUser(ctx, req)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.UserNotFoundError
		}
		return nil, err
	}

	return &domain.GetUserResponse{User: user}, nil
}

func (c *core) DeleteUser(ctx context.Context, req *domain.DeleteUserRequest) (*domain.DeleteUserResponse, error) {
	_, err := c.GetUser(ctx, &domain.GetUserRequest{UserId: req.UserID})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.DeleteUser(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.DeleteUserResponse{}, nil
}
