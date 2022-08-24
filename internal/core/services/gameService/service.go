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

func (srv *service) Create(name string, size uint, bombs uint) (game.Game, error) {
	if bombs >= size*size {
		return game.Game{}, errors.New("the number of bombs is invalid")
	}

	game := game.NewGame(uuid.New().String(), name, size, bombs)

	if err := srv.gamesRepository.Save(game); err != nil {
		return game.Game{}, errors.New("create game into repository has failed")
	}

	return game, nil
}
