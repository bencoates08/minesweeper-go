package game

import (
	board "minesweeper-go/internal/core/domain/board"
)

const (
	StateInProgress = "in progress"
	StateWon        = "won"
	StateLost       = "lost"
)

type Game struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	State          string              `json:"state"`
	BoardSettings  board.BoardSettings `json:"board_settings"`
	Board          board.Board         `json:"board"`
	PlayerView     [][]string          `json:"player_view"`
	CellsRemaining int                 `json:"cells_remaining"`
}

func NewGame(id string, name string, height int, width int, bombs int) Game {
	boardSettings := board.NewBoardSettings(height, width, bombs)
	board, _ := board.NewBoard(boardSettings)
	return Game{
		ID:             id,
		Name:           name,
		State:          StateInProgress,
		BoardSettings:  boardSettings,
		Board:          board,
		PlayerView:     board.GetVisibleBoard(),
		CellsRemaining: board.CellsRemaining,
	}
}

func (g *Game) Reveal(row int, col int) {
	position := g.Board.GetPosition(row-1, col-1)

	err := g.Board.Reveal(position)
	g.PlayerView = g.Board.GetVisibleBoard()
	if err != nil {
		if err.Error() == "bomb hit" {
			g.State = StateLost
			return
		}
	}

	g.CellsRemaining = g.Board.CellsRemaining
	if g.CellsRemaining == 0 {
		g.State = StateWon
	}
}
