package errcore

var (
	PlatformNotFoundError = &CoreError{
		Message: "Platform not found.",
		Code:    "platform.not_found",
		Type:    NotFoundType,
	}

	PlatformsNotFoundError = &CoreError{
		Message: "Platforms not found.",
		Code:    "platforms.not_found",
		Type:    NotFoundType,
	}
)
