package game_test

import (
	games "minesweeper-go/internal/core/domain/game"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	Describe("NewGame", func() {
		When("the game settings are valid", func() {
			It("creates a new game", func() {
				// Act
				game, err := games.NewGame("id", "name", 10, 10, 10)

				// Assert
				Expect(err).To(BeNil())

				Expect(game.ID).To(Equal("id"))
				Expect(game.Name).To(Equal("name"))
				Expect(game.State).To(Equal(games.StateInProgress))
				Expect(game.CellsRemaining).To(Equal(90))

				Expect(game.BoardSettings.Height).To(Equal(10))
				Expect(game.BoardSettings.Width).To(Equal(10))
				Expect(game.BoardSettings.Bombs).To(Equal(10))
				Expect(game.Board.CellsRemaining).To(Equal(90))

				Expect(game.PlayerView).To(HaveLen(10))
				Expect(game.PlayerView[0]).To(HaveLen(10))

				Expect(game.Board).NotTo(BeNil())
			})
		})

		When("the game settings are invalid and fails to create a board", func() {
			Context("height provided is below 1", func() {
				It("returns an error", func() {
					// Act
					_, err := games.NewGame("id", "name", 0, 10, 10)

					// Assert
					Expect(err).To(MatchError("unable to create new game: board dimensions must be greater than 0"))
				})
			})

			Context("width provided is below 1", func() {
				It("returns an error", func() {
					// Act
					_, err := games.NewGame("id", "name", 10, 0, 10)

					// Assert
					Expect(err).To(MatchError("unable to create new game: board dimensions must be greater than 0"))
				})
			})

			Context("number of mines provided is below 1", func() {
				It("returns an error", func() {
					// Act
					_, err := games.NewGame("id", "name", 10, 10, 0)

					// Assert
					Expect(err).To(MatchError("unable to create new game: number of bombs must be greater than 0"))
				})
			})

			Context("number of mines provided is greater than the number of cells", func() {
				It("returns an error", func() {
					// Act
					_, err := games.NewGame("id", "name", 10, 10, 1000)

					// Assert
					Expect(err).To(MatchError("unable to create new game: too many bombs"))
				})
			})
		})
	})

	Describe("Reveal", func() {
		When("a valid cell is revealed", func() {
			When("the cell is a normal cell", func() {
				It("reveals the cell", func() {
					/**
					Game board:
					{"1", "X", "X", "3", "X", "X"},
					{"1", "2", "2", "3", "X", "4"},
					{"1", "1", "-", "1", "2", "X"},
					{"X", "1", "-", "-", "1", "1"},
					{"2", "2", "1", "-", "-", "-"},
					{"1", "X", "1", "-", "-", "-"},
					*/

					// Arrange
					game, _ := games.NewGame("id", "name", 6, 6, 8, 1)

					// Assert
					Expect(game.PlayerView[0][0]).To(Equal("H"))

					// Act
					err := game.Reveal(1, 1)

					// Assert
					Expect(err).ToNot(HaveOccurred())
					Expect(game.State).To(Equal(games.StateInProgress))
					Expect(game.CellsRemaining).To(Equal(27))
					Expect(game.PlayerView[0][0]).To(Equal("1"))
				})
			})

			When("the cell revealed is a bomb", func() {
				It("reveals the cell and ends the game", func() {
					/**
					Game board:
					{"1", "X"},
					*/

					// Arrange
					game, _ := games.NewGame("id", "name", 1, 2, 1, 1)

					// Act
					err := game.Reveal(1, 2)

					// Assert
					Expect(err).ToNot(HaveOccurred())
					Expect(game.PlayerView[0][1]).To(Equal("X"))
					Expect(game.State).To(Equal(games.StateLost))
				})
			})

			When("the cell is the final cell", func() {
				It("reveals the cell and ends the game", func() {
					/**
					Game board:
					{"1", "X"},
					*/

					// Arrange
					game, _ := games.NewGame("id", "name", 1, 2, 1, 1)

					// Act
					err := game.Reveal(1, 1)

					// Assert
					Expect(err).ToNot(HaveOccurred())
					Expect(game.PlayerView[0][0]).To(Equal("1"))
					Expect(game.CellsRemaining).To(Equal(0))
					Expect(game.State).To(Equal(games.StateWon))
				})
			})
		})

		When("there is an error revealing the cell", func() {
			It("returns an error", func() {
				// Arrange
				game, _ := games.NewGame("id", "name", 1, 2, 1, 1)

				// Act
				err := game.Reveal(10, 10)

				// Assert
				Expect(err).To(MatchError("unable to reveal cell: row index out of bounds"))
			})
		})

		When("the game is no longer in progress", func() {
			It("returns an error", func() {
				/**
				Game board:
				{"1", "X"},
				*/

				// Arrange
				game, _ := games.NewGame("id", "name", 1, 2, 1, 1)
				game.Reveal(1, 1)

				// Act
				err := game.Reveal(1, 2)

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("game is no longer in progress, the game is won"))
			})
		})
	})
})
