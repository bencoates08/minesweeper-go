package main

import (
	"minesweeper-go/internal/core/services/gameService"
	"minesweeper-go/internal/repositories/gameRepository"
	// "minesweeper-go/internal/core/services/gameService"
	// "minesweeper-go/internal/handlers/gameHandler"
	// "minesweeper-go/internal/repositories/gameRepository"
	// "github.com/gin-gonic/gin"
)

func main() {
	gameRepository := gameRepository.NewMemKVS()
	gameService := gameService.New(gameRepository)

	gameID, _ := gameService.Create("test", 20, 30, 80)

	gameService.Reveal(gameID, 0, 0)
	gameService.Reveal(gameID, 0, 4)
	gameService.Reveal(gameID, 1, 2)
	gameService.Reveal(gameID, 15, 15)

	game, _ := gameService.Get(gameID)

	game.Board.PrintRevealedBoard()
	game.Board.Print()

	// boardDisplay.FlagPosition(Position{Row: 1, Col: 3, Val: "F"})

	// gameRepository := gameRepository.NewMemKVS()
	// gameService := gameService.New(gameRepository)
	// gameHandler := gameHandler.NewHTTPHandler(gameService)

	// router := gin.New()
	// router.GET("/games/:id", gameHandler.Get)
	// router.POST("/games", gameHandler.Create)

	// router.Run(":8080")
}
