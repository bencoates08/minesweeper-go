package board

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	CELL_EMPTY  = "-"
	CELL_BOMB   = "X"
	CELL_HIDDEN = "H"
)

type Board struct {
	BoardState     [][]Position `json:"board_state"`
	CellsRemaining int          `json:"cells_remaining"`
}

// Create a new Minesweeper board used to track game state and display the board.
func NewBoard(boardSettings BoardSettings, seed int64) (Board, error) {
	rand.Seed(seed)
	// Check provided board dimensions are valid
	if boardSettings.Height < 1 || boardSettings.Width < 1 {
		return Board{}, errors.New("board dimensions must be greater than 0")
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
	err := board.addBombs(boardSettings.Bombs)
	if err != nil {
		return Board{}, err
	}
	board.calculateBoardThreatLevels()

	return board, nil
}

// Randomly disperse bombs within BoardState
func (b *Board) addBombs(bombs int) error {
	// Check number of bombs is valid
	if len(b.BoardState)*len(b.BoardState[0])-bombs < 1 {
		return errors.New("too many bombs")
	}

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

		var err error
		boardPositionsRemaining, err = removeElement(boardPositionsRemaining, randPosIndex)
		if err != nil {
			return fmt.Errorf("error removing element from boardPositionsRemaining: %v", err)
		}

		b.BoardState[randPos.Row][randPos.Col].Val = CELL_BOMB
	}

	return nil
}

// Remove an element from a slice.
func removeElement(slice []Position, index int) ([]Position, error) {
	if index < 0 || index > len(slice)-1 {
		return slice, errors.New("index out of bounds")
	}
	return append(slice[:index], slice[index+1:]...), nil
}

// Get Position object from board coordinates.
func (b Board) getPosition(row int, col int) (Position, error) {
	if row < 0 || row > len(b.BoardState)-1 {
		return Position{}, errors.New("row index out of bounds")
	}
	if col < 0 || col > len(b.BoardState[row])-1 {
		return Position{}, errors.New("column index out of bounds")
	}

	return b.BoardState[row][col], nil
}

// Get neighbouring positions of a given position on the board.
func (b Board) getNeighbouringPositions(position Position) []*Position {
	// row_limit and column_limit define board boundaries to prevent out of bounds errors
	row_limit := len(b.BoardState) - 1
	column_limit := len(b.BoardState[0]) - 1
	neighbouringPositions := make([]*Position, 0)

	for i := max(0, position.Row-1); i <= min(position.Row+1, row_limit); i++ {
		for j := max(0, position.Col-1); j <= min(position.Col+1, column_limit); j++ {
			if i != position.Row || j != position.Col {
				neighbouringPositions = append(neighbouringPositions, &b.BoardState[i][j])
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
func (b *Board) Reveal(row int, col int) error {
	position, err := b.getPosition(row, col)
	if err != nil {
		return err
	}

	if position.Val == CELL_BOMB {
		b.BoardState[position.Row][position.Col].Visible = true
		return errors.New("bomb hit")
	}

	if position.Visible {
		return errors.New("cell already revealed")
	}

	b.BoardState[position.Row][position.Col].Visible = true
	b.CellsRemaining--

	neighbouringPositions := b.getNeighbouringPositions(position)
	for _, neighbouringPosition := range neighbouringPositions {

		if !neighbouringPosition.Visible {
			if neighbouringPosition.Val == CELL_EMPTY {
				b.Reveal(neighbouringPosition.Row, neighbouringPosition.Col)
			} else if position.Val == CELL_EMPTY {
				b.BoardState[neighbouringPosition.Row][neighbouringPosition.Col].Visible = true
				b.CellsRemaining--
			}
		}
	}

	return nil
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
