package board

import (
	"errors"
	"math/rand"
	"strconv"
)

const (
	CELL_EMPTY  = "-"
	CELL_BOMB   = "X"
	CELL_HIDDEN = "H"
)

type Position struct {
	Row int
	Col int
	Val string
}

type Board struct {
	BoardValues     [][]string `json:"board_values"`
	BoardVisibility [][]bool   `json:"board_visibility"`
	CellsRemaining  int        `json:"cells_remaining"`
}

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
	board := Board{BoardValues: boardValues, BoardVisibility: boardVisibility}
	
	board.CellsRemaining = boardSettings.Height*boardSettings.Width - boardSettings.Bombs

	// Initialize bombs and threat level to boardValues
	board.addBombs(boardSettings.Bombs)
	board.calculateBoardThreatLevels()

	return board
}

// Randomly disperse bombs within boardValues
func (b *Board) addBombs(bombs int) {

	// Initialize boardPositionsRemaining, a slice to track remaining positions to be filled with bombs.
	// helps prevent randomly picking cells until a free cell is found.
	boardPositionsRemaining := make([]Position, len(b.BoardValues)*len(b.BoardValues[0]))
	for row := range b.BoardValues {
		for col := range b.BoardValues[0] {
			boardPositionsRemaining[row*len(b.BoardValues[0])+col] = Position{Row: row, Col: col}
		}
	}

	// Randomly select an index from board positions remaining and place a bomb there then remove
	// position from remaining positions.
	for i := 0; i < bombs; i++ {
		randPosIndex := rand.Intn(len(boardPositionsRemaining))
		randPos := boardPositionsRemaining[randPosIndex]
		boardPositionsRemaining = removeElement(boardPositionsRemaining, randPosIndex)
		b.BoardValues[randPos.Row][randPos.Col] = CELL_BOMB
	}
}

// Remove an elemnt from a slice.
func removeElement(slice []Position, index int) []Position {
	return append(slice[:index], slice[index+1:]...)
}

// Get Position object from board coordinates.
func (b Board) GetPosition(row int, col int) Position {
	return Position{Row: row, Col: col, Val: b.BoardValues[row][col]}
}

// Get neighbouring positions of a given position on the board.
func (b Board) getNeighbouringPositions(position Position) []Position {
	// row_limit and column_limit define board boundaries to prevent out of bounds errors
	row_limit := len(b.BoardValues) - 1
	column_limit := len(b.BoardValues[0]) - 1
	neighbouringPositions := make([]Position, 0)

	for i := max(0, position.Row-1); i <= min(position.Row+1, row_limit); i++ {
		for j := max(0, position.Col-1); j <= min(position.Col+1, column_limit); j++ {
			if i != position.Row || j != position.Col {
				neighbouringPositions = append(neighbouringPositions, Position{Row: i, Col: j, Val: b.BoardValues[i][j]})
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
	for row := range b.BoardValues {
		for col := range b.BoardValues[0] {
			if b.BoardValues[row][col] == CELL_EMPTY {
				threatLevel := b.getThreatLevel(Position{Row: row, Col: col})

				if threatLevel > 0 {
					b.BoardValues[row][col] = strconv.Itoa(threatLevel)
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
func (b *Board) Reveal(position Position) error {
	if position.Val == CELL_BOMB {
		b.BoardVisibility[position.Row][position.Col] = true
		return errors.New("bomb hit")
	}

	if !b.isCellRevealed(position) {
		b.BoardVisibility[position.Row][position.Col] = true
		b.CellsRemaining--

		neighbouringPositions := b.getNeighbouringPositions(position)
		for _, neighbouringPosition := range neighbouringPositions {

			if !b.isCellRevealed(neighbouringPosition) {
				if neighbouringPosition.Val == CELL_EMPTY {
					b.Reveal(neighbouringPosition)
				} else if position.Val == CELL_EMPTY {
					b.BoardVisibility[neighbouringPosition.Row][neighbouringPosition.Col] = true
					b.CellsRemaining--
				}
			}
		}
	}
	return nil
}

// Checks if a cell is revealed.
func (b Board) isCellRevealed(position Position) bool {
	return b.BoardVisibility[position.Row][position.Col]
}

// Gets the current view of the board.
func (b Board) GetVisibleBoard() [][]string {

	visibleBoard := make([][]string, len(b.BoardValues))
	for row := range visibleBoard {
		visibleBoard[row] = make([]string, len(b.BoardValues[0]))
	}

	for row := range visibleBoard {
		for col := range visibleBoard[0] {
			if b.BoardVisibility[row][col] {
				visibleBoard[row][col] = b.BoardValues[row][col]
			} else {
				visibleBoard[row][col] = CELL_HIDDEN
			}
		}
	}

	return visibleBoard
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
