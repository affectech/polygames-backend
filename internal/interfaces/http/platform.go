package http

import (
	"net/http"
	"polygames/internal/domain"
)

// CreatePlatform godoc
// @Summary Create a new platform (NOT IN PRODUCTION)
// @Description Create a new platform with provided details
// @Tags platforms
// @Accept json
// @Produce json
// @Param request body domain.CreatePlatformRequest true "Create platform request body"
// @Success 201 {object} domain.Platform "Successfully created platform"
// @Router /api/v1/platforms [post]
func (s *server) CreatePlatform(w http.ResponseWriter, r *http.Request) {
	var req domain.CreatePlatformRequest
	Platform, err := s.core.CreatePlatform(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusCreated, Platform, w)
}

// GetPlatform godoc
// @Summary Get a platform by ID (NOT IN PRODUCTION)
// @Description Retrieve a platform by its ID
// @Tags platforms
// @Accept json
// @Produce json
// @Param platform_id path integer true "Platform ID"
// @Success 200 {object} domain.Platform "Successfully retrieved platform"
// @Router /api/v1/platforms/{platform_id} [get]
func (s *server) GetPlatform(w http.ResponseWriter, r *http.Request) {
	req := domain.GetPlatformRequest{PlatformID: s.parseParamInt64("platform_id", r)}
	Platform, err := s.core.GetPlatform(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, Platform, w)
}

// UpdatePlatform godoc
// @Summary Update a platform by ID (NOT IN PRODUCTION)
// @Description Update details of a platform by its ID
// @Tags platforms
// @Accept json
// @Produce json
// @Param platform_id path integer true "Platform ID"
// @Param request body domain.UpdatePlatformRequest true "Update platform request body"
// @Success 200 {object} domain.Platform "Successfully updated platform"
// @Router /api/v1/platforms/{platform_id} [put]
func (s *server) UpdatePlatform(w http.ResponseWriter, r *http.Request) {
	var req domain.UpdatePlatformRequest
	if err := s.readJSON(&req, r); err != nil {
		s.sendError(err, w)
		return
	}

	req.ID = s.parseParamInt64("platform_id", r)

	Platform, err := s.core.UpdatePlatform(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, Platform, w)
}

// DeletePlatform godoc
// @Summary Delete a platform by ID (NOT IN PRODUCTION)
// @Description Delete a platform by its ID
// @Tags platforms
// @Accept json
// @Produce json
// @Param platform_id path integer true "Platform ID"
// @Success 200 "Successfully deleted platform"
// @Router /api/v1/platforms/{platform_id} [delete]
func (s *server) DeletePlatform(w http.ResponseWriter, r *http.Request) {
	req := domain.DeletePlatformRequest{PlatformID: s.parseParamInt64("platform_id", r)}

	_, err := s.core.DeletePlatform(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, nil, w)
}
