package board_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"minesweeper-go/internal/core/domain/board"
)

var _ = Describe("Board", func() {
	PREDEFINED_BOARD_STATE := [][]string{
		{"1", "X", "X", "3", "X", "X"},
		{"1", "2", "2", "3", "X", "4"},
		{"1", "1", "-", "1", "2", "X"},
		{"X", "1", "-", "-", "1", "1"},
		{"2", "2", "1", "-", "-", "-"},
		{"1", "X", "1", "-", "-", "-"},
	}

	Describe("NewBoard", func() {
		When("the board settings are valid", func() {
			It("should return a board with the correct dimensions", func() {
				// Arrange
				boardSettings := board.NewBoardSettings(6, 6, 8)

				// Act
				board, err := board.NewBoard(boardSettings, 1)

				// Assert
				Expect(err).To(BeNil())
				Expect(len(board.BoardState)).To(Equal(6))
				Expect(len(board.BoardState[0])).To(Equal(6))
				Expect(board.CellsRemaining).To(Equal(28))
				checkBoardState(board, PREDEFINED_BOARD_STATE)
			})
		})

		When("the board settings are invalid", func() {
			Context("height provided is below 1", func() {
				It("should return an error", func() {
					// Arrange
					boardSettings := board.NewBoardSettings(0, 6, 8)

					// Act
					_, err := board.NewBoard(boardSettings, 1)

					// Assert
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("board dimensions must be greater than 0"))
				})
			})

			Context("width provided is below 1", func() {
				It("should return an error", func() {
					// Arrange
					boardSettings := board.NewBoardSettings(6, 0, 8)

					// Act
					_, err := board.NewBoard(boardSettings, 1)

					// Assert
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("board dimensions must be greater than 0"))
				})
			})

			Context("number of bombs is more than available cells", func() {
				It("should return an error", func() {
					// Arrange
					boardSettings := board.NewBoardSettings(6, 6, 50)

					// Act
					_, err := board.NewBoard(boardSettings, 1)

					// Assert
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("too many bombs"))
				})
			})
		})
	})

	Describe("Reveal", func() {
		var currentBoard board.Board

		BeforeEach(func() {
			// Arrange
			boardSettings := board.NewBoardSettings(6, 6, 8)
			currentBoard, _ = board.NewBoard(boardSettings, 1)
		})

		When("the cell is a number and not adjacent to empty cell", func() {
			It("should reveal only that cell and cells remaining reduces", func() {
				// Act
				err := currentBoard.Reveal(0, 0)

				// Assert
				Expect(err).ToNot(HaveOccurred())
				Expect(currentBoard.BoardState[0][0].Visible).To(BeTrue())
				validateAllPositions(currentBoard, positionIsNotVisible, currentBoard.BoardState[0][0])
				Expect(currentBoard.CellsRemaining).To(Equal(27))
			})
		})

		When("the cell is a number and adjacent to empty cell", func() {
			It("should reveal all cells adjacent to adjacent empty cells and cells remaining reduces", func() {
				// Act
				err := currentBoard.Reveal(3, 5)

				// Assert
				Expect(err).ToNot(HaveOccurred())
				Expect(currentBoard.BoardState[3][5].Visible).To(BeTrue())
				Expect(currentBoard.CellsRemaining).To(Equal(7))
			})
		})
		When("the cell is empty", func() {
			It("should reveal all cells adjacent to adjacent empty cells and cells remaining reduces", func() {
				// Act
				err := currentBoard.Reveal(5, 5)

				// Assert
				Expect(err).ToNot(HaveOccurred())
				Expect(currentBoard.BoardState[5][5].Visible).To(BeTrue())
				Expect(currentBoard.CellsRemaining).To(Equal(7))
			})
		})

		When("the cell is already revealed", func() {
			It("should return an error", func() {
				// Act
				err := currentBoard.Reveal(0, 0)

				// Assert
				Expect(err).ToNot(HaveOccurred())

				// Act
				err = currentBoard.Reveal(0, 0)

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("cell already revealed"))
			})
		})

		When("the cell is a bomb", func() {
			It("should return a game over error", func() {
				// Act
				err := currentBoard.Reveal(0, 4)

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("bomb hit"))
			})
		})

		When("the cell is out of bounds", func() {})
	})

	Describe("GetVisibleBoard", func() {
		var currentBoard board.Board

		BeforeEach(func() {
			// Arrange
			boardSettings := board.NewBoardSettings(6, 6, 8)
			currentBoard, _ = board.NewBoard(boardSettings, 1)
		})

		When("the board is in its initial state", func() {
			It("should return a board with all cells hidden", func() {
				// Arrange
				expectedBoard := [][]string{
					{"H", "H", "H", "H", "H", "H"},
					{"H", "H", "H", "H", "H", "H"},
					{"H", "H", "H", "H", "H", "H"},
					{"H", "H", "H", "H", "H", "H"},
					{"H", "H", "H", "H", "H", "H"},
					{"H", "H", "H", "H", "H", "H"},
				}

				// Act
				visibleBoard := currentBoard.GetVisibleBoard()

				// Assert
				Expect(visibleBoard).To(Equal(expectedBoard))
			})
		})

		When("the board has some cells revealed", func() {
			It("should return a board with all cells hidden", func() {
				// Arrange
				expectedBoard := [][]string{
					{"1", "H", "H", "H", "H", "H"},
					{"H", "2", "2", "3", "H", "H"},
					{"H", "1", "-", "1", "2", "H"},
					{"H", "1", "-", "-", "1", "1"},
					{"H", "2", "1", "-", "-", "-"},
					{"H", "H", "1", "-", "-", "-"},
				}

				// Act
				currentBoard.Reveal(0, 0)
				currentBoard.Reveal(5, 5)
				visibleBoard := currentBoard.GetVisibleBoard()

				// Assert
				Expect(visibleBoard).To(Equal(expectedBoard))
			})
		})
	})
})

//
// Helper functions

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

func validateAllPositions(board board.Board, validateFn func(board.Position) (bool, error), ignorePositions ...board.Position) {
	for _, row := range board.BoardState {
		for _, cell := range row {
			if !containsPosition(ignorePositions, cell) {
				valid, err := validateFn(cell)
				if !valid {
					Fail(fmt.Sprintf("Expected cell at (%d, %d) to be valid: %v", cell.Row+1, cell.Col+1, err))
				}
			}
		}
	}
}

func containsPosition(positions []board.Position, position board.Position) bool {
	for _, pos := range positions {
		if pos == position {
			return true
		}
	}
	return false
}

//
// Validation functions

func positionIsNotVisible(pos board.Position) (bool, error) {
	if pos.Visible {
		return false, errors.New("position was found to be visible")
	}
	return true, nil
}
