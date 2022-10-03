package board_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"minesweeper-go/internal/core/domain/board"
)

var _ = Describe("Board", func() {
	PREDEFINED_BOARD_STATE := [][]string{
		{"1", "1", "2", "X", "1"},
		{"1", "X", "2", "1", "1"},
		{"3", "3", "2", "1", "1"},
		{"X", "X", "1", "1", "X"},
		{"2", "2", "1", "1", "1"},
	}

	Describe("Initialising a board", func() {
		When("the board settings are valid", func() {
			It("should return a board with the correct dimensions", func() {
				// Arrange
				boardSettings := board.NewBoardSettings(5, 5, 5)

				// Act
				board, err := board.NewBoard(boardSettings)

				// Assert
				Expect(err).To(BeNil())
				Expect(len(board.BoardState)).To(Equal(5))
				Expect(len(board.BoardState[0])).To(Equal(5))
				Expect(board.CellsRemaining).To(Equal(20))
				checkBoardState(board, PREDEFINED_BOARD_STATE)
			})
		})

		When("the board settings are invalid", func() {
			Context("height provided is below 1", func() {
				It("should return an error", func() {
					// Arrange
					boardSettings := board.NewBoardSettings(0, 5, 5)
	
					// Act
					_, err := board.NewBoard(boardSettings)
	
					// Assert
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("board dimensions must be greater than 0"))
				})
			})

			Context("width provided is below 1", func() {
				It("should return an error", func() {
					// Arrange
					boardSettings := board.NewBoardSettings(5, 0, 5)
	
					// Act
					_, err := board.NewBoard(boardSettings)
	
					// Assert
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("board dimensions must be greater than 0"))
				})
			})

			Context("number of bombs is more than available cells", func() {
				It("should return an error", func() {
					// Arrange
					boardSettings := board.NewBoardSettings(5, 5, 50)
	
					// Act
					_, err := board.NewBoard(boardSettings)
	
					// Assert
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("too many bombs"))
				})
			})
		})
	})
})

func checkBoardState(board board.Board, expectedState [][]string) {
	for i, row := range expectedState {
		for j, cell := range row {
			if board.BoardState[i][j].Val != cell {
				Fail(fmt.Sprintf("Expected cell at (%d, %d) to be %s, but got %s", i+1, j+1, cell, board.BoardState[i][j].Val))
			}
			if board.BoardState[i][j].Visible {
				Fail(fmt.Sprintf("Expected cell at (%d, %d) to be unrevealed", i+1, j+1))
			}
		}
	}
}
