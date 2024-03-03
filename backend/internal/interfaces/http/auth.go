package http

import (
	_ "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"
	"net/http"
	_ "polygames/cmd/docs"
	"polygames/internal/domain"
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
