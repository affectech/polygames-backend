package games

import "context"

type Games interface {
	Start() error
	Stop(ctx context.Context) error
}
