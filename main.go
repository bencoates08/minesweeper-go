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
	board := NewBoard(boardSettings)

	board.PrintRevealedBoard()
	board.Print()

	// input := ""
	// for input != "q" {
	// 	println("Enter row: ")
	// 	fmt.Scanln(&input)
	// 	println("Enter col: ")
	// 	fmt.Scanln(&input)

	// 	row, _ := strconv.Atoi(input)
	// 	col, _ := strconv.Atoi(input)
	// 	boardDisplay.Reveal(boardState.GetPosition(col-1, row-1), boardState)
	// 	boardDisplay.Print(boardState)
	// }

	board.Reveal(board.GetPosition(0, 0))
	board.Print()
	println()

	board.Reveal(board.GetPosition(0, 4))
	board.Print()
	println()

	board.Reveal(board.GetPosition(1, 2))
	board.Print()
	println()

	board.Reveal(board.GetPosition(15, 15))
	board.Print()
	println()

	// boardDisplay.FlagPosition(Position{Row: 1, Col: 3, Val: "F"})

	// gameRepository := gameRepository.NewMemKVS()
	// gameService := gameService.New(gameRepository)
	// gameHandler := gameHandler.NewHTTPHandler(gameService)

	// router := gin.New()
	// router.GET("/games/:id", gameHandler.Get)
	// router.POST("/games", gameHandler.Create)

	// router.Run(":8080")
}
