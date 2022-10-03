package gameHandler

import "minesweeper-go/internal/core/domain/game"

type CreateRequest struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Bombs  int    `json:"bombs"`
}

type RevealRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type GameResponse struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	State          string     `json:"state"`
	CellsRemaining int        `json:"cells_remaining"`
	Board          [][]string `json:"board"`
}

func BuildGameResponse(model game.Game) GameResponse {
	return GameResponse{
		ID:             model.ID,
		Name:           model.Name,
		State:          model.State,
		CellsRemaining: model.CellsRemaining,
		Board:          model.PlayerView,
	}
}
