package domain

import board "minesweeper-go/internal/core/domain/Board"

type Game struct {
	ID            string              `json:"id"`
	Name          string              `json:"name"`
	State         string              `json:"state"`
	BoardSettings board.BoardSettings `json:"board_settings"`
	Board         board.BoardState    `json:"board"`
}
