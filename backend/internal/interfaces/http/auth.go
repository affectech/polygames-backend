package http

import (
	"fmt"
	_ "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"
	"net/http"
	_ "polygames/cmd/docs"
	"polygames/internal/core"
	"polygames/internal/domain"
	"time"
)

// SignIn godoc
// @Summary Sign in
// @Description Sign in with provided credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param request body domain.SignInRequest true "Sign in request body"
// @Success 200 {object} domain.SignInResponse
// @Router /api/v1/auth/sign-in [post]
func (s *server) SignIn(w http.ResponseWriter, r *http.Request) {
	var req domain.SignInRequest
	if err := s.readJSON(&req, r); err != nil {
		s.sendError(err, w)
		return
	}

	response, err := s.core.SignIn(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    response.SessionID,
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(core.TTL),
	})

	s.sendJSON(http.StatusOK, response, w)
}

// SignUp godoc
// @Summary Sign up
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body domain.SignUpRequest true "Sign up request body"
// @Success 200 {object} domain.SignUpResponse
// @Router /api/v1/auth/sign-up [post]
func (s *server) SignUp(w http.ResponseWriter, r *http.Request) {
	var req domain.SignUpRequest
	if err := s.readJSON(&req, r); err != nil {
		s.sendError(err, w)
		return
	}

	response, err := s.core.SignUp(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, response, w)
}

func (s *server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "session not found", http.StatusUnauthorized)
			return
		}

		sess, err := core.GetSession(cookie.Value)
		if err != nil {
			http.Error(w, fmt.Sprintf("user session not found(session id: %s), err: %v", cookie.Value, err), http.StatusUnauthorized)
			return
		}

		token := r.Header.Get("X-CSRF-Token")
		if token != sess.CSRFToken {
			http.Error(w, "invalid csrf token("+token+")", http.StatusUnauthorized)
			return
		}

		var getUserRequest *domain.GetUserRequest
		getUserRequest.UserId = sess.UserID

		user, err := s.core.GetUser(r.Context(), getUserRequest)
		if err != nil {
			http.Error(w, "user not found", http.StatusUnauthorized)
			return
		}
		if user.User == nil {
			http.Error(w, "no data on user", http.StatusUnauthorized)
			return
		}

		ctx := core.NewContext(r.Context(), &domain.User{
			ID:       user.User.ID,
			Username: user.User.Username,
			Email:    user.User.Email,
			Role:     user.User.Role,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
