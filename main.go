package main

import (
	. "minesweeper-go/internal/core/domain/board"
	// "minesweeper-go/internal/core/services/gameService"
	// "minesweeper-go/internal/handlers/gameHandler"
	// "minesweeper-go/internal/repositories/gameRepository"
	// "github.com/gin-gonic/gin"
)

func main() {
	boardSettings := BoardSettings{Height: 20, Width: 30, Bombs: 80}
	boardState := NewBoardState(boardSettings)
	boardDisplay := NewBoardDisplay(boardSettings)

	boardState.Print()

	println()

	boardDisplay.GetPositionsFromReveal(boardState.GetPosition(0, 11), boardState, make([]Position, 0))
	boardDisplay.Print(boardState)

	println()

	boardDisplay.GetPositionsFromReveal(boardState.GetPosition(0, 1), boardState, make([]Position, 0))
	boardDisplay.Print(boardState)

	println()

	boardDisplay.GetPositionsFromReveal(boardState.GetPosition(15, 14), boardState, make([]Position, 0))
	boardDisplay.Print(boardState)

	// boardDisplay.FlagPosition(Position{Row: 1, Col: 3, Val: "F"})

	// gameRepository := gameRepository.NewMemKVS()
	// gameService := gameService.New(gameRepository)
	// gameHandler := gameHandler.NewHTTPHandler(gameService)

	// router := gin.New()
	// router.GET("/games/:id", gameHandler.Get)
	// router.POST("/games", gameHandler.Create)

	// router.Run(":8080")
}
