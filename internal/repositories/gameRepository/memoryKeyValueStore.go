package gameRepository

import (
	"encoding/json"
	"errors"
	"minesweeper-go/internal/core/domain/game"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (game.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		currentGame := game.Game{}
		err := json.Unmarshal(value, &currentGame)
		if err != nil {
			return game.Game{}, errors.New("fail to get value from kvs")
		}

		return currentGame, nil
	}

	return game.Game{}, errors.New("game not found in kvs")
}

func (repo *memkvs) Save(currentGame game.Game) error {
	value, err := json.Marshal(currentGame)
	if err != nil {
		return errors.New("fail to save value into kvs")
	}

	repo.kvs[currentGame.ID] = value

	return nil
}
