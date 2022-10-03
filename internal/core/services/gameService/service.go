package gameService

import (
	"fmt"
	"minesweeper-go/internal/core/domain/game"
	"minesweeper-go/internal/core/ports"

	"github.com/google/uuid"
)

type service struct {
	gamesRepository ports.GamesRepository
}

func New(gamesRepository ports.GamesRepository) *service {
	return &service{
		gamesRepository: gamesRepository,
	}
}

func (srv *service) Get(id string) (game.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		return game, fmt.Errorf("unable to get a game with the given id (%v): %v", id, err)
	}

	return game, nil
}

func (srv *service) Create(name string, height int, width int, bombs int) (game.Game, error) {
	newGame, err := game.NewGame(uuid.New().String(), name, height, width, bombs)
	if err != nil {
		return game.Game{}, err
	}

	err = srv.gamesRepository.Save(newGame)
	if err != nil {
		return game.Game{}, fmt.Errorf("unable to save new game to repository: %v", err)
	}

	return newGame, nil
}

func (srv *service) Reveal(id string, row int, col int) (game.Game, error) {
	currentGame, err := srv.Get(id)
	if err != nil {
		return currentGame, err
	}

	err = currentGame.Reveal(row, col)
	if err != nil {
		return currentGame, err
	}

	err = srv.gamesRepository.Save(currentGame)
	if err != nil {
		return currentGame, fmt.Errorf("unable to save the game to the repository: %v", err)
	}

	return currentGame, nil
}
