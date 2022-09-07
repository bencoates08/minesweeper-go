package gameService

import (
	"errors"
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
		return game, errors.New("get game from repository has failed")
	}

	return game, nil
}

func (srv *service) Create(name string, height int, width int, bombs int) (game.Game, error) {
	if bombs >= height*width-9 {
		return game.Game{}, errors.New("the number of bombs is invalid")
	}

	newGame := game.NewGame(uuid.New().String(), name, height, width, bombs)

	err := srv.gamesRepository.Save(newGame)
	if err != nil {
		return game.Game{}, errors.New("create game into repository has failed")
	}

	return newGame, nil
}
