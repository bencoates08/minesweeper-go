package game

import (
	. "minesweeper-go/internal/core/domain/Board"

	"github.com/google/uuid"
)

type Game struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	State         string        `json:"state"`
	BoardSettings BoardSettings `json:"board_settings"`
	Board         Board         `json:"board"`
}

func NewGame(id string, name string, height int, width int, bombs int) Game {
	boardSettings := NewBoardSettings(height, width, bombs)
	return Game{
		ID:            uuid.New().String(),
		Name:          name,
		State:         "started",
		BoardSettings: boardSettings,
		Board:         NewBoard(boardSettings),
	}
}

func (g *Game) Reveal(row int, col int) {
	position := g.Board.GetPosition(row, col)
	g.Board.Reveal(position)
}
