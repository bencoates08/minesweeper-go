package board

import (
	"fmt"
	"math/rand"
	"strconv"
)

// TODO: handle flagging
const (
	CELL_EMPTY   = "-"
	CELL_BOMB    = "X"
	CELL_HIDDEN  = "H"
	CELL_FLAGGED = "F"
)

type Position struct {
	Row int
	Col int
	Val string
}

type Board struct {
	boardValues     [][]string
	boardVisibility [][]bool
}

// TODO: ask question about returning
// Create a new Minesweeper board used to track game state and display the board.
func NewBoard(boardSettings BoardSettings) Board {
	// Initialize boardVisibility
	boardVisibility := make([][]bool, boardSettings.Height)
	for row := range boardVisibility {
		boardVisibility[row] = make([]bool, boardSettings.Width)
	}

	// Initialize empty boardValues
	boardValues := make([][]string, boardSettings.Height)
	for row := range boardValues {
		boardValues[row] = make([]string, boardSettings.Width)
	}

	for row := range boardValues {
		for col := range boardValues[0] {
			boardValues[row][col] = CELL_EMPTY
		}
	}

	// Initialize board
	board := Board{boardValues: boardValues, boardVisibility: boardVisibility}

	// Initialize bombs and threat level to boardValues
	board.addBombs(boardSettings.Bombs)
	board.calculateBoardThreatLevels()

	return board
}

// Randomly disperse bombs within boardValues
func (b *Board) addBombs(bombs int) {

	// Initialize boardPositionsRemaining, a slice to track remaining positions to be filled with bombs.
	// helps prevent randomly picking cells until a free cell is found.
	boardPositionsRemaining := make([]Position, len(b.boardValues)*len(b.boardValues[0]))
	for row := range b.boardValues {
		for col := range b.boardValues[0] {
			boardPositionsRemaining[row*len(b.boardValues[0])+col] = Position{Row: row, Col: col}
		}
	}

	// Randomly select an index from board positions remaining and place a bomb there then remove
	// position from remaining positions.
	for i := 0; i < bombs; i++ {
		randPosIndex := rand.Intn(len(boardPositionsRemaining))
		randPos := boardPositionsRemaining[randPosIndex]
		boardPositionsRemaining = removeElement(boardPositionsRemaining, randPosIndex)
		b.boardValues[randPos.Row][randPos.Col] = CELL_BOMB
	}
}

// Remove an elemnt from a slice.
func removeElement(slice []Position, index int) []Position {
	return append(slice[:index], slice[index+1:]...)
}

// Get Position object from board coordinates.
func (b Board) GetPosition(row int, col int) Position {
	return Position{Row: row, Col: col, Val: b.boardValues[row][col]}
}

// Get neighbouring positions of a given position on the board.
func (b Board) getNeighbouringPositions(position Position) []Position {
	// row_limit and column_limit define board boundaries to prevent out of bounds errors
	row_limit := len(b.boardValues) - 1
	column_limit := len(b.boardValues[0]) - 1
	neighbouringPositions := make([]Position, 0)

	for i := max(0, position.Row-1); i <= min(position.Row+1, row_limit); i++ {
		for j := max(0, position.Col-1); j <= min(position.Col+1, column_limit); j++ {
			if i != position.Row || j != position.Col {
				neighbouringPositions = append(neighbouringPositions, Position{Row: i, Col: j, Val: b.boardValues[i][j]})
			}
		}
	}

	return neighbouringPositions
}

// Get threat level of a given position on the board.
//
// Threat level is the number of bombs surrounding a cell including diagonally adjacent cells.
func (b Board) getThreatLevel(position Position) int {
	neighbouringPositions := b.getNeighbouringPositions(position)
	threatLevel := 0

	for _, neighbouringPosition := range neighbouringPositions {
		if neighbouringPosition.Val == CELL_BOMB {
			threatLevel++
		}
	}

	return threatLevel
}

// Calculate board threat levels.
//
// Calculate threat level for each cell on the board and store in boardValues.
func (b *Board) calculateBoardThreatLevels() {
	for row := range b.boardValues {
		for col := range b.boardValues[0] {
			if b.boardValues[row][col] == CELL_EMPTY {
				threatLevel := b.getThreatLevel(Position{Row: row, Col: col})

				if threatLevel > 0 {
					b.boardValues[row][col] = strconv.Itoa(threatLevel)
				}
			}
		}
	}
}

// Reveals board cells from selected board position.
//
// If the selected cell or one of its adjacent cells has 0 threat level then all
// cells adjacent to the zero threat level cell are revealed. All adjacent 0 threat
// level cells are revealed recursively i.e. a chunk of the board is revealed.
func (b *Board) Reveal(position Position) {
	if position.Val == CELL_BOMB {
		return
	}

	if !b.isCellRevealed(position) {
		b.boardVisibility[position.Row][position.Col] = true

		neighbouringPositions := b.getNeighbouringPositions(position)
		for _, neighbouringPosition := range neighbouringPositions {

			if !b.isCellRevealed(neighbouringPosition) {
				if neighbouringPosition.Val == CELL_EMPTY {
					b.Reveal(neighbouringPosition)
				} else if position.Val == CELL_EMPTY {
					b.boardVisibility[neighbouringPosition.Row][neighbouringPosition.Col] = true
				}
			}
		}
	}
}

// Checks if a cell is revealed.
func (b Board) isCellRevealed(position Position) bool {
	return b.boardVisibility[position.Row][position.Col]
}

// Prints the current view of the board.
func (b Board) Print() {
	for i, row := range b.boardVisibility {
		for j := range row {
			if b.boardVisibility[i][j] {
				fmt.Print(CELL_HIDDEN)
				print(" ")
			} else {
				fmt.Print(b.boardValues[i][j])
				print(" ")
			}
		}
		println()
	}
}

// Prints the fully revealed board.
func (b Board) PrintRevealedBoard() {
	for _, row := range b.boardValues {
		fmt.Println(row)
	}
}

// Get max of two integers.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Get min of two integers.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
