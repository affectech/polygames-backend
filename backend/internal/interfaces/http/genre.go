package http

import (
	"net/http"
	"polygames/internal/domain"
)

func (s *server) CreateGenre(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateGenreRequest
	genre, err := s.core.CreateGenre(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusCreated, genre, w)
}

func (s *server) GetGenre(w http.ResponseWriter, r *http.Request) {
	req := domain.GetGenreRequest{GenreID: s.parseParamInt64("genre_id", r)}
	Genre, err := s.core.GetGenre(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, Genre, w)
}

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

func (s *server) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	req := domain.DeleteGenreRequest{GenreID: s.parseParamInt64("genre_id", r)}

	_, err := s.core.DeleteGenre(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, nil, w)
}
