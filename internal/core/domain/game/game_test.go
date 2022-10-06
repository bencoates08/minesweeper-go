package game_test

import (
	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
	// "minesweeper-go/internal/core/domain/game"
)

var _ = Describe("Game", func() {
	Describe("NewGame", func() {
		When("the game settings are valid", func() {})

		When("the game settings are invalid and fails to create a board", func() {
			Context("height provided is below 1", func() {})

			Context("width provided is below 1", func() {})

			Context("number of mines provided is below 1", func() {})
			
			Context("number of mines provided is greater than the number of cells", func() {})
		})
	})

	Describe("Reveal", func() {
		When("a valid cell is revealed", func() {
			When("the cell is a normal cell", func() {})

			When("the cell is the final cell", func() {})

			When("the cell revealed is a bomb", func() {})
		})
		
		When("there is an error revealing the cell", func() {})

		When("the game is no longer in progress", func() {})
	})
})
