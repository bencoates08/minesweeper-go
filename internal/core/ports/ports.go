package ports

import "minesweeper-go/internal/core/domain/game"

type GamesRepository interface {
	Get(id string) (game.Game, error)
	Save(game.Game) error
}

type GamesService interface {
	Get(id string) (game.Game, error)
	Create(name string, height int, width int, bombs int) (game.Game, error)
	Reveal(id string, row int, col int) (game.Game, error)
}
