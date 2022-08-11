package board

import (
	"fmt"
	"math/rand"
	"strconv"
)

// TODO: Seperate 2D array to track which cells are revealed and Flaged.
type BoardState [][]string

// TODO: add positon to its own file
type Position struct {
	Row int
	Col int
	Val string
}

const (
	CELL_EMPTY = "-"
	CELL_BOMB  = "X"
)

func NewBoardState(boardSettings BoardSettings) BoardState {
	boardState := make(BoardState, boardSettings.Height)
	for row := range boardState {
		boardState[row] = make([]string, boardSettings.Width)
	}

	for row := range boardState {
		for col := range boardState[0] {
			boardState[row][col] = CELL_EMPTY
		}
	}

	boardState.addBombs(boardSettings.Bombs)
	boardState.calculateBoardNumbers()

	return boardState
}

func (boardState BoardState) addBombs(bombs int) {
	boardPositionsRemaining := make([]Position, len(boardState)*len(boardState[0]))
	for row := range boardState {
		for col := range boardState[0] {
			boardPositionsRemaining[row*len(boardState[0])+col] = Position{Row: row, Col: col}
		}
	}

	for i := 0; i < bombs; i++ {
		randPosIndex := rand.Intn(len(boardPositionsRemaining))
		randPos := boardPositionsRemaining[randPosIndex]
		boardPositionsRemaining = removeIndex(boardPositionsRemaining, randPosIndex)

		boardState[randPos.Row][randPos.Col] = CELL_BOMB
	}
}

func removeIndex(slice []Position, index int) []Position {
	return append(slice[:index], slice[index+1:]...)
}

func (boardState BoardState) GetPosition(row int, col int) Position {
	return Position{Row: row, Col: col, Val: boardState[row][col]}
}

func (boardState BoardState) getNeighbouringPositions(position Position) []Position {
	row_limit := len(boardState) - 1
	column_limit := len(boardState[0]) - 1
	neighbouringPositions := make([]Position, 0)

	for i := max(0, position.Row-1); i <= min(position.Row+1, row_limit); i++ {
		for j := max(0, position.Col-1); j <= min(position.Col+1, column_limit); j++ {
			if i != position.Row || j != position.Col {
				neighbouringPositions = append(neighbouringPositions, Position{Row: i, Col: j, Val: boardState[i][j]})
			}
		}
	}

	return neighbouringPositions
}

func (boardState BoardState) getThreatLevel(position Position) int {
	neighbouringPositions := boardState.getNeighbouringPositions(position)
	threatLevel := 0

	for _, neighbouringPosition := range neighbouringPositions {
		if neighbouringPosition.Val == CELL_BOMB {
			threatLevel++
		}
	}

	return threatLevel
}

func (boardState BoardState) calculateBoardNumbers() {
	for row := range boardState {
		for col := range boardState[0] {
			if boardState[row][col] == CELL_EMPTY {
				threatLevel := boardState.getThreatLevel(Position{Row: row, Col: col})

				if threatLevel > 0 {
					boardState[row][col] = strconv.Itoa(threatLevel)
				}
			}
		}
	}
}

func (boardState BoardState) Print() {
	for _, row := range boardState {
		fmt.Println(row)
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
