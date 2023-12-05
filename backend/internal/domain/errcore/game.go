package errcore

var (
	GameNotFoundError = &CoreError{
		Message: "Game not found.",
		Code:    "game.not_found",
		Type:    NotFoundType,
	}

	GamesNotFoundError = &CoreError{
		Message: "Games not found.",
		Code:    "games.not_found",
		Type:    NotFoundType,
	}
)
