package board

import (
	"errors"
	"math/rand"
	"strconv"
)

const (
	CELL_EMPTY = "-"
	CELL_BOMB  = "X"
	CELL_HIDDEN = "H"
)

type Board struct {
	BoardState     [][]Position `json:"board_state"`
	CellsRemaining int          `json:"cells_remaining"`
}

// Create a new Minesweeper board used to track game state and display the board.
func NewBoard(boardSettings BoardSettings) (Board, error) {
	// Check provided board dimensions are valid
	if boardSettings.Height < 1 || boardSettings.Width < 1 {
		return Board{}, errors.New("board dimensions must be greater than 0")
	} else if boardSettings.Height*boardSettings.Width - boardSettings.Bombs < 1 {
		return Board{}, errors.New("too many bombs")
	}

	// Initialize empty boardState
	boardState := make([][]Position, boardSettings.Height)
	for row := range boardState {
		boardState[row] = make([]Position, boardSettings.Width)
	}

	for row := range boardState {
		for col := range boardState[0] {
			boardState[row][col] = Position{Row: row, Col: col, Val: CELL_EMPTY, Visible: false, Flagged: false}
		}
	}

	// Initialize board
	cellsRemaining := boardSettings.Height*boardSettings.Width - boardSettings.Bombs
	board := Board{BoardState: boardState, CellsRemaining: cellsRemaining}

	// Initialize bombs and threat level to BoardState
	board.addBombs(boardSettings.Bombs)
	board.calculateBoardThreatLevels()

	return board, nil
}

// Randomly disperse bombs within BoardState
func (b *Board) addBombs(bombs int) {

	// Initialize boardPositionsRemaining, a slice to track remaining positions to be filled with bombs.
	// helps prevent randomly picking cells until a free cell is found.
	boardPositionsRemaining := make([]Position, len(b.BoardState)*len(b.BoardState[0]))
	for row := range b.BoardState {
		for col := range b.BoardState[0] {
			boardPositionsRemaining[row*len(b.BoardState[0])+col] = b.BoardState[row][col]
		}
	}

	// Randomly select an index from board positions remaining and place a bomb there then remove
	// position from remaining positions.
	for i := 0; i < bombs; i++ {
		randPosIndex := rand.Intn(len(boardPositionsRemaining))
		randPos := boardPositionsRemaining[randPosIndex]
		boardPositionsRemaining = removeElement(boardPositionsRemaining, randPosIndex)
		b.BoardState[randPos.Row][randPos.Col].Val = CELL_BOMB
	}
}

// Remove an elemnt from a slice.
func removeElement(slice []Position, index int) []Position {
	return append(slice[:index], slice[index+1:]...)
}

// Get Position object from board coordinates.
func (b Board) GetPosition(row int, col int) Position {
	return b.BoardState[row][col]
}

// Get neighbouring positions of a given position on the board.
func (b Board) getNeighbouringPositions(position Position) []Position {
	// row_limit and column_limit define board boundaries to prevent out of bounds errors
	row_limit := len(b.BoardState) - 1
	column_limit := len(b.BoardState[0]) - 1
	neighbouringPositions := make([]Position, 0)

	for i := max(0, position.Row-1); i <= min(position.Row+1, row_limit); i++ {
		for j := max(0, position.Col-1); j <= min(position.Col+1, column_limit); j++ {
			if i != position.Row || j != position.Col {
				neighbouringPositions = append(neighbouringPositions, b.BoardState[i][j])
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
// Calculate threat level for each cell on the board and store in BoardState.
func (b *Board) calculateBoardThreatLevels() {
	for row := range b.BoardState {
		for col := range b.BoardState[0] {
			if b.BoardState[row][col].Val == CELL_EMPTY {
				threatLevel := b.getThreatLevel(b.BoardState[row][col])

				if threatLevel > 0 {
					b.BoardState[row][col].Val = strconv.Itoa(threatLevel)
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
		b.BoardState[position.Row][position.Col].Visible = true
		return errors.New("bomb hit")
	}

	if !b.isCellRevealed(position) {
		b.BoardState[position.Row][position.Col].Visible = true
		b.CellsRemaining--

		neighbouringPositions := b.getNeighbouringPositions(position)
		for _, neighbouringPosition := range neighbouringPositions {

			if !b.isCellRevealed(neighbouringPosition) {
				if neighbouringPosition.Val == CELL_EMPTY {
					b.Reveal(neighbouringPosition)
				} else if position.Val == CELL_EMPTY {
					b.BoardState[neighbouringPosition.Row][neighbouringPosition.Col].Visible = true
					b.CellsRemaining--
				}
			}
		}
	}
	return nil
}

// Checks if a cell is revealed.
func (b Board) isCellRevealed(position Position) bool {
	return b.BoardState[position.Row][position.Col].Visible
}

// Gets the current view of the board.
func (b Board) GetVisibleBoard() [][]string {
	visibleBoard := make([][]string, len(b.BoardState))
	for row := range visibleBoard {
		visibleBoard[row] = make([]string, len(b.BoardState[0]))
	}

	for row := range visibleBoard {
		for col := range visibleBoard[0] {
			if b.BoardState[row][col].Visible {
				visibleBoard[row][col] = b.BoardState[row][col].Val
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
