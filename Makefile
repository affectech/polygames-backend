.PHONY:run
run:
	go run cmd/main.go
migrate:
	go run cmd/migrate.go

.PHONY:build-w
build:
	go build -o build/polygames.exe cmd/main.go

.PHONY:docs
docs:
	swag init -g ./internal/interfaces/http/auth.go -o ./cmd/docs
	swag init -g ./internal/interfaces/http/game.go -o ./cmd/docs
	swag init -g ./internal/interfaces/http/genre.go -o ./cmd/docs
	swag init -g ./internal/interfaces/http/user.go -o ./cmd/docs
	swag init -g ./internal/interfaces/http/platform.go -o ./cmd/docs
