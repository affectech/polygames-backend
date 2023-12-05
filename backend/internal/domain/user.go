package domain

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	_ Gender = iota
	Male
	Female
)

const (
	_ UserRole = iota
	UserRoleGlobalAdmin
	UserRoleAdmin
	UserRoleModerator
	UserRoleUser
)

type (
	Gender   int8
	UserRole int16
)
type (
	User struct {
		ID        int64      `json:"id"`
		Username  string     `json:"username"`
		Email     string     `json:"email"`
		Password  string     `json:"-"`
		Gender    Gender     `json:"gender"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
		Role      UserRole   `json:"role"`
	}

	CreateUserRequest struct {
		ID       int64    `json:"id"`
		Username string   `json:"username"`
		Email    string   `json:"email"`
		Password string   `json:"-"`
		Gender   Gender   `json:"gender"`
		Role     UserRole `json:"role"`
	}
	CreateUserResponse struct {
		User *User `json:"data"`
	}

	GetUserRequest struct {
		UserId int64 `json:"-"`
	}
	GetUserResponse struct {
		User *User `json:"data"`
	}
	GetUsersRequest struct {
		UserId int64 `json:"-"`
	}
	GetUsersResponse struct {
		Users []User `json:"data"`
	}
	UpdateUserRequest struct {
		UserID   int64     `json:"user_id"`
		Username *string   `json:"username"`
		Email    *string   `json:"email"`
		Password *string   `json:"password"`
		Gender   Gender    `json:"gender"`
		Role     *UserRole `json:"role"`
	}
	UpdateUserResponse struct {
		User *User `json:"data"`
	}

	DeleteUserRequest struct {
		UserID int64 `json:"user_id"`
	}
	DeleteUserResponse struct{}
)

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
