package ports

import "minesweeper-go/internal/core/domain/game"

type GamesRepository interface {
	Get(id string) (game.Game, error)
	Save(game.Game) error
}

type GamesService interface {
	Get(id string) (game.Game, error)
	Create(name string, size uint, bombs uint) (game.Game, error)
	Reveal(id string, row uint, col uint) (game.Game, error)
}
