package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"polygames/internal/domain"
	"strconv"
	"strings"
)

func (s *server) CreateGame(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateGameRequest
	//UserID
	userIDString := r.FormValue("user_id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing user_id: %w", err), w)
		return
	}
	req.UserID = int64(userID)
	//TeamID
	teamIDString := r.FormValue("team_id")
	teamID, err := strconv.Atoi(teamIDString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing team_id: %w", err), w)
		return
	}
	req.TeamID = int64(teamID)
	//Description
	req.Description = strings.TrimSpace(r.FormValue("description"))
	//GenreID
	genreIDString := r.FormValue("genre_id")
	genreID, err := strconv.Atoi(genreIDString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing team_id: %w", err), w)
		return
	}
	req.GenreID = int64(genreID)
	// Получаем файл из формы запроса
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, fmt.Sprintf("parsing image file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Создаем новый файл на сервере для сохранения изображения
	imagePath := "./images/" + header.Filename
	out, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("error saving image: %s", err), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Копируем данные из загруженного файла в новый файл на сервере
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("error coping image: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Image uploaded successfully.")
	return

	game, err := s.core.CreateGame(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusCreated, game, w)
}

func (s *server) GetGame(w http.ResponseWriter, r *http.Request) {
	req := domain.GetGameRequest{GameID: s.parseParamInt64("game_id", r)}
	game, err := s.core.GetGame(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, game, w)
}

func (s *server) UpdateGame(w http.ResponseWriter, r *http.Request) {
	var req domain.UpdateGameRequest
	if err := s.readJSON(&req, r); err != nil {
		s.sendError(err, w)
		return
	}

	req.ID = s.parseParamInt64("game_id", r)

	game, err := s.core.UpdateGame(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, game, w)
}

func (s *server) DeleteGame(w http.ResponseWriter, r *http.Request) {
	req := domain.DeleteGameRequest{GameID: s.parseParamInt64("game_id", r)}

	_, err := s.core.DeleteGame(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, nil, w)
}
