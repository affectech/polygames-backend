package http

import (
	"net/http"
	"polygames/internal/domain"
)

func (s *server) CreatePlatform(w http.ResponseWriter, r *http.Request) {
	var req domain.CreatePlatformRequest
	Platform, err := s.core.CreatePlatform(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusCreated, Platform, w)
}

func (s *server) GetPlatform(w http.ResponseWriter, r *http.Request) {
	req := domain.GetPlatformRequest{PlatformID: s.parseParamInt64("platform_id", r)}
	Platform, err := s.core.GetPlatform(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, Platform, w)
}

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

func (s *server) DeletePlatform(w http.ResponseWriter, r *http.Request) {
	req := domain.DeletePlatformRequest{PlatformID: s.parseParamInt64("platform_id", r)}

	_, err := s.core.DeletePlatform(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, nil, w)
}
