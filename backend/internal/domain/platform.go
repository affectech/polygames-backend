package domain

type (
	Platform struct {
		ID   int64  `json:"id"`
		Name string `json:"description"`
	}
	CreatePlatformRequest struct {
		ID   int64  `json:"id"`
		Name string `json:"description"`
	}

	CreatePlatformResponse struct {
		Platform *Platform `json:"data"`
	}
	GetPlatformRequest struct {
		PlatformID int64 `json:"platform_id"`
	}
	GetPlatformResponse struct {
		Platform *Platform `json:"data"`
	}
	UpdatePlatformRequest struct {
		ID   int64  `json:"id"`
		Name string `json:"description"`
	}
	UpdatePlatformResponse struct {
		Platform *Platform `json:"data"`
	}
	DeletePlatformRequest struct {
		PlatformID int64 `json:"platform_id"`
	}
	DeletePlatformResponse struct {
	}
)
