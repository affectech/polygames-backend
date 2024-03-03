package http

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"polygames/internal/domain"
	"strconv"
	"strings"
)

// CreateGame godoc
// @Summary Create a new game
// @Description Create a new game with provided details
// @Tags games
// @Accept json
// @Produce json
// @Param user_id query integer true "User ID"
// @Param team_id query integer true "Team ID"
// @Param description query string true "Description"
// @Param genre_id query integer true "Genre ID"
// @Param image formData file true "Game Image"
// @Param file formData file true "Game File"
// @Success 201 {object} domain.Game "Successfully created game"
// @Router /api/v1/games [post]
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
	//Image
	// Получаем файл из формы запроса
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, fmt.Sprintf("parsing image file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	// Создаем новый файл на сервере для сохранения изображения
	imagePath := "./images/" + header.Filename
	out, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("error saving image: %s", err), http.StatusInternalServerError)
		return
	}
	defer func(out *os.File) {
		err = out.Close()
		if err != nil {

		}
	}(out)

	// Копируем данные из загруженного файла в новый файл на сервере
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("error coping image: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "image uploaded successfully.")

	game, err := s.core.CreateGame(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusCreated, game, w)
	//File
	// Получаем файл из формы запроса
	file, header, err = r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("parsing game file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	// Создаем новый файл на сервере для сохранения файла игры
	gamePath := "./images/" + header.Filename
	out, err = os.Create(gamePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("error saving game file: %s", err), http.StatusInternalServerError)
		return
	}
	defer func(out *os.File) {
		err = out.Close()
		if err != nil {

		}
	}(out)

	// Копируем данные из загруженного файла в новый файл на сервере
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("error coping game file: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "game file uploaded successfully.")

	game, err = s.core.CreateGame(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusCreated, game, w)
}

// GetGame godoc
// @Summary Get a game by ID
// @Description Retrieve a game by its ID
// @Tags games
// @Accept json
// @Produce json
// @Param game_id path integer true "Game ID"
// @Success 200 {object} domain.Game "Successfully retrieved game"
// @Router /api/v1/games/{game_id} [get]
func (s *server) GetGame(w http.ResponseWriter, r *http.Request) {
	req := domain.GetGameRequest{GameID: s.parseParamInt64("game_id", r)}
	game, err := s.core.GetGame(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, game, w)
}

// UpdateGame godoc
// @Summary Update a game by ID
// @Description Update details of a game by its ID
// @Tags games
// @Accept json
// @Produce json
// @Param game_id path integer true "Game ID"
// @Param request body domain.UpdateGameRequest true "Update game request body"
// @Success 200 {object} domain.Game "Successfully updated game"
// @Router /api/v1/games/{game_id} [put]
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

// DeleteGame godoc
// @Summary Delete a game by ID
// @Description Delete a game by its ID
// @Tags games
// @Accept json
// @Produce json
// @Param game_id path integer true "Game ID"
// @Success 200 "Successfully deleted game"
// @Router /api/v1/games/{game_id} [delete]
func (s *server) DeleteGame(w http.ResponseWriter, r *http.Request) {
	req := domain.DeleteGameRequest{GameID: s.parseParamInt64("game_id", r)}

	_, err := s.core.DeleteGame(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, nil, w)
}
