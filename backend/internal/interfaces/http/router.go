package http

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	_ "polygames/cmd/docs"
)

func swaggerHandler(res http.ResponseWriter, req *http.Request) {
	httpSwagger.WrapHandler(res, req)
}

func (s *server) initRouter() http.Handler {
	r := s.router

	if os.Getenv("ENV") == "dev" {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	r.Get("/swagger/*", swaggerHandler)

	r.Post("/api/v1/auth/sign-in", s.SignIn)
	r.Post(`/api/v1/auth/sign-up`, s.SignUp)

	r.Get(`/api/v1/users/{user_id}`, s.GetUser)

	r.Get(`/api/v1/games/{game_id}`, s.GetGame)
	r.Post(`/api/v1/games`, s.CreateGame)
	r.Put(`/api/v1/games/{game_id}`, s.UpdateGame)
	r.Delete(`/api/v1/games/{game_id}`, s.DeleteGame)

	return r
}
