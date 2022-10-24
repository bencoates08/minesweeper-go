package gameService_test

import (
	"minesweeper-go/internal/core/ports"
	. "minesweeper-go/internal/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	ctrl := gomock.NewController(GinkgoT())
  defer ctrl.Finish()
	
	var MockGamesRepository ports.GamesRepository

	BeforeEach(func() {
		MockGamesRepository = NewMockGamesRepository(ctrl)
	})

	Describe("Get", func() {
		When("the game exists in the database", func() {
			It("should return a game", func() {})
		})

		When("the game does not exist in the database", func() {
			It("should return an error", func() {})
		})
	})

	Describe("Create", func() {

	})

	Describe("Save", func() {
		
	})

	Describe("Reveal", func() {
		When("the game exists in the database", func() {
			It("should reveal a cell", func() {})
		})

		When("the game does not exist in the database", func() {
			It("should return an error", func() {})
		})
	})
})
