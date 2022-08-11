package gameRepository

import (
	"encoding/json"
	"errors"
	"minesweeper-go/internal/core/domain"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (domain.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return domain.Game{}, errors.New("fail to get value from kvs")
		}

		return game, nil
	}

	return domain.Game{}, errors.New("game not found in kvs")
}
