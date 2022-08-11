package gameService

import (
	"errors"
	"minesweeper-go/internal/core/domain"
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

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New("get game from repository has failed")
	}

	return game, nil
}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New("the number of bombs is invalid")
	}

	game := domain.NewGame(uuid.New().String(), name, size, bombs)

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New("create game into repository has failed")
	}

	return game, nil
}
