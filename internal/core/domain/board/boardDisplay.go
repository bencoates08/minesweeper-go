package board

import "fmt"

type BoardDisplay [][]bool

const (
	CELL_HIDDEN  = "H"
	CELL_FLAGGED = "F"
)

func NewBoardDisplay(boardSettings BoardSettings) BoardDisplay {
	boardDisplay := make(BoardDisplay, boardSettings.Height)
	for row := range boardDisplay {
		boardDisplay[row] = make([]bool, boardSettings.Width)
	}

	return boardDisplay
}

func (boardDisplay BoardDisplay) GetPositionsFromReveal(position Position, boardState BoardState, revealedPositions []Position) []Position {
	if position.Val == CELL_BOMB {
		return append(revealedPositions, position)
	}

	if !visitedCell(boardDisplay, position) {
		revealedPositions = append(revealedPositions, position)
		boardDisplay[position.Row][position.Col] = true

		neighbouringPositions := boardState.getNeighbouringPositions(position)
		for _, neighbouringPosition := range neighbouringPositions {

			if !visitedCell(boardDisplay, neighbouringPosition) {
				if boardState[neighbouringPosition.Row][neighbouringPosition.Col] == CELL_EMPTY {
					boardDisplay.GetPositionsFromReveal(neighbouringPosition, boardState, revealedPositions)
				} else if position.Val == CELL_EMPTY {
					revealedPositions = append(revealedPositions, neighbouringPosition)
					boardDisplay[neighbouringPosition.Row][neighbouringPosition.Col] = true
				}
			}

		}

	}

	return revealedPositions
}

func visitedCell(visitedPositions [][]bool, position Position) bool {
	return visitedPositions[position.Row][position.Col]
}

func (boardDisplay BoardDisplay) Print(boardState BoardState) {
	for i, row := range boardDisplay {
		for j := range row {
			if boardDisplay[i][j] == false {
				fmt.Print(CELL_HIDDEN)
			} else {
				fmt.Print(boardState[i][j])
			}
		}
		println()
	}
}

func (boardDisplay BoardDisplay) PrintVis() {
	for _, row := range boardDisplay {
		fmt.Println(row)
	}
}

// fuc (boardDisplay BoardDisplay) FlagPosition(positions Position) error {
// 	f boardDisplay[positions.Row][positions.Col] == CELL_HIDDEN {
// 		boardDisplay[positions.Row][positions.Col] = CELL_FLAGGED
// 		return nil
// 	} else {
// 		return mt.Errorf("Position already revealed")
// 	}
// }
