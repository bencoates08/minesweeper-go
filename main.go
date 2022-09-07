package main

import (
	"minesweeper-go/internal/core/services/gameService"
	"minesweeper-go/internal/handlers/gameHandler"
	"minesweeper-go/internal/repositories/gameRepository"

	"github.com/gin-gonic/gin"
)

func main() {
	gameRepository := gameRepository.NewMemKVS()
	gameService := gameService.New(gameRepository)
	gameHandler := gameHandler.NewHTTPHandler(gameService)

	router := gin.New()
	router.GET("/games/:id", gameHandler.Get)
	router.POST("/games", gameHandler.Create)
	router.POST("/games/:id/reveal", gameHandler.Reveal)

	router.Run(":8080")
}
