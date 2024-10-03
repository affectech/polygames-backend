package domain

import "time"

type (
	AccessTokenMetadata struct {
		Id       int64    `json:"id"`
		Username string   `json:"username"`
		Email    string   `json:"email"`
		Role     UserRole `json:"role"`
	}
	RefreshTokenMetadata struct {
		ID int64 `json:"id"`
	}

	SignInRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	SignInResponse struct {
		SessionID    string `json:"-"`
		CSRFToken    string `json:"csrfToken"`
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}

	SignUpRequest struct {
		Username    string    `json:"username"`
		Email       string    `json:"email"`
		Password    string    `json:"password"`
		Gender      string    `json:"gender"`
		DateOfBirth time.Time `json:"date_of_birth"`
	}
	SignUpResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
)
