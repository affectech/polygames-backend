package http

import (
	"net/http"
	"polygames/internal/domain"
)

// GetUser godoc
// @Summary Get a user by ID
// @Description Retrieve a user by its ID
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path integer true "User ID"
// @Success 200 {object} domain.User "Successfully retrieved user"
// @Router /api/v1/users/{user_id} [get]
func (s *server) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := s.parseParamInt64("user_id", r)

	user, err := s.core.GetUser(r.Context(), &domain.GetUserRequest{UserId: userId})
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, user, w)
}
