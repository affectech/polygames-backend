package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"polygames/internal/domain"
	"polygames/internal/repository/database"
)

func (d *driver) GetUser(ctx context.Context, req *domain.GetUserRequest) (*domain.User, error) {
	row := d.conn.QueryRow(ctx, `select id, username, email, gender, created_at, updated_at, role
                                     from users
                                     where id = $1`, req.UserId)

	var user domain.User
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Gender,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrObjectNotFound
		}
		return nil, fmt.Errorf("scanning user by credentials: %w", err)
	}

	return &user, nil
}

func (d *driver) GetUserByLogin(ctx context.Context, login string) (*domain.User, error) {
	var user domain.User
	err := d.conn.QueryRow(ctx, `select id, username, email, password, gender, created_at, updated_at, role
                                     from users
                                     where lower(username) = $1 or lower(email) = $1`, login).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Gender,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrObjectNotFound
		}
		return nil, fmt.Errorf("scanning user by credentials: %w", err)
	}

	return &user, nil
}

func (d *driver) DeleteUser(ctx context.Context, req *domain.DeleteUserRequest) error {
	err := d.conn.QueryRow(ctx, `delete 
									 from users 
									 where id=$1`, req.UserID)
	if err != nil {
		return fmt.Errorf("delete error %w", err)
	}
	return nil
}

func (d *driver) SignUp(ctx context.Context, req *domain.SignUpRequest) (int64, error) {
	var id int64
	err := d.conn.QueryRow(ctx, `insert into users(username, email, password, role) 
									 values($1, $2, $3, $4) returning id`,
		req.Username,
		req.Email,
		req.Password,
		domain.UserRoleUser,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("inserting user for sign up: %w", err)
	}

	return id, nil
}
