package ports

import (
	"context"
	"minesweeper-go/internal/core/domain/game"
)

type GamesRepository interface {
	Get(context.Context, string) (game.Game, error)
	Save(context.Context, game.Game) error
}

type GamesService interface {
	Get(ctx context.Context, id string) (game.Game, error)
	Create(ctx context.Context, name string, height int, width int, bombs int) (game.Game, error)
	Reveal(ctx context.Context, id string, row int, col int) (game.Game, error)
}
