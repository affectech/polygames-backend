package http

import (
	"net/http"
	"polygames/internal/domain"
)

// CreateGenre godoc
// @Summary Create a new genre (NOT IN PRODUCTION)
// @Description Create a new genre with provided details
// @Tags genres
// @Accept json
// @Produce json
// @Param request body domain.CreateGenreRequest true "Create genre request body"
// @Success 201 {object} domain.Genre "Successfully created genre"
// @Router /api/v1/genres [post]
func (s *server) CreateGenre(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateGenreRequest
	genre, err := s.core.CreateGenre(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusCreated, genre, w)
}

// GetGenre godoc
// @Summary Get a genre by ID (NOT IN PRODUCTION)
// @Description Retrieve a genre by its ID
// @Tags genres
// @Accept json
// @Produce json
// @Param genre_id path integer true "Genre ID"
// @Success 200 {object} domain.Genre "Successfully retrieved genre"
// @Router /api/v1/genres/{genre_id} [get]
func (s *server) GetGenre(w http.ResponseWriter, r *http.Request) {
	req := domain.GetGenreRequest{GenreID: s.parseParamInt64("genre_id", r)}
	Genre, err := s.core.GetGenre(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, Genre, w)
}

// UpdateGenre godoc
// @Summary Update a genre by ID (NOT IN PRODUCTION)
// @Description Update details of a genre by its ID
// @Tags genres
// @Accept json
// @Produce json
// @Param genre_id path integer true "Genre ID"
// @Param request body domain.UpdateGenreRequest true "Update genre request body"
// @Success 200 {object} domain.Genre "Successfully updated genre"
// @Router /api/v1/genres/{genre_id} [put]
func (s *server) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	var req domain.UpdateGenreRequest
	if err := s.readJSON(&req, r); err != nil {
		s.sendError(err, w)
		return
	}

	req.ID = s.parseParamInt64("genre_id", r)

	Genre, err := s.core.UpdateGenre(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, Genre, w)
}

// DeleteGenre godoc
// @Summary Delete a genre by ID (NOT IN PRODUCTION)
// @Description Delete a genre by its ID
// @Tags genres
// @Accept json
// @Produce json
// @Param genre_id path integer true "Genre ID"
// @Success 200 "Successfully deleted genre"
// @Router /api/v1/genres/{genre_id} [delete]
func (s *server) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	req := domain.DeleteGenreRequest{GenreID: s.parseParamInt64("genre_id", r)}

	_, err := s.core.DeleteGenre(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, nil, w)
}
