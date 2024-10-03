package domain

type (
	Genre struct {
		ID   int64  `json:"id"`
		Name string `json:"description"`
	}
	CreateGenreRequest struct {
		ID   int64  `json:"id"`
		Name string `json:"description"`
	}

	CreateGenreResponse struct {
		Genre *Genre `json:"data"`
	}
	GetGenreRequest struct {
		GenreID int64 `json:"genre_id"`
	}
	GetGenreResponse struct {
		Genre *Genre `json:"data"`
	}
	UpdateGenreRequest struct {
		ID   int64  `json:"id"`
		Name string `json:"description"`
	}
	UpdateGenreResponse struct {
		Genre *Genre `json:"data"`
	}
	DeleteGenreRequest struct {
		GenreID int64 `json:"genre_id"`
	}
	DeleteGenreResponse struct {
	}
)
