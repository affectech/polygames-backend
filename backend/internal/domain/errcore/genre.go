package errcore

var (
	GenreNotFoundError = &CoreError{
		Message: "Genre not found.",
		Code:    "genre.not_found",
		Type:    NotFoundType,
	}

	GenresNotFoundError = &CoreError{
		Message: "Genres not found.",
		Code:    "genres.not_found",
		Type:    NotFoundType,
	}
)
