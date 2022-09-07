package game

import (
	board "minesweeper-go/internal/core/domain/Board"
)

type Game struct {
	ID            string              `json:"id"`
	Name          string              `json:"name"`
	State         string              `json:"state"`
	BoardSettings board.BoardSettings `json:"board_settings"`
	Board         board.Board         `json:"board"`
}

func NewGame(id string, name string, height int, width int, bombs int) Game {
	boardSettings := board.NewBoardSettings(height, width, bombs)
	return Game{
		ID:            id,
		Name:          name,
		State:         "started",
		BoardSettings: boardSettings,
		Board:         board.NewBoard(boardSettings),
	}
}

func (g *Game) Reveal(row int, col int) {
	position := g.Board.GetPosition(row, col)
	g.Board.Reveal(position)
}
