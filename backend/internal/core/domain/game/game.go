package game

import (
	"fmt"
	board "minesweeper-go/internal/core/domain/board"
	"time"
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

func NewGame(id string, name string, height int, width int, bombs int, seed ...int64) (Game, error) {
	var boardSettings board.BoardSettings
	if len(seed) > 0 {
		boardSettings = board.NewBoardSettings(height, width, bombs, seed[0])
	} else {
		boardSettings = board.NewBoardSettings(height, width, bombs, time.Now().UnixNano())
	}

	board, err := board.NewBoard(boardSettings)
	if err != nil {
		return Game{}, fmt.Errorf("unable to create new game: %v", err)
	}

	return Game{
			ID:             id,
			Name:           name,
			State:          StateInProgress,
			BoardSettings:  boardSettings,
			Board:          board,
			PlayerView:     board.GetVisibleBoard(),
			CellsRemaining: board.CellsRemaining,
		},
		nil
}

func (g *Game) Reveal(row int, col int) error {
	if g.State != StateInProgress {
		return fmt.Errorf("game is no longer in progress, the game is %v", g.State)
	}

	err := g.Board.Reveal(row-1, col-1)
	g.PlayerView = g.Board.GetVisibleBoard()
	if err != nil {
		if err.Error() == "bomb hit" {
			g.State = StateLost
			return nil
		}
		return fmt.Errorf("unable to reveal cell: %v", err)
	}

	g.CellsRemaining = g.Board.CellsRemaining
	if g.CellsRemaining == 0 {
		g.State = StateWon
	}

	return nil
}
