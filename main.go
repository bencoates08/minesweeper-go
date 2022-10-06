package main

import (
	"database/sql"
	"fmt"
	"minesweeper-go/config"
	"minesweeper-go/internal/core/services/gameService"
	"minesweeper-go/internal/handlers/gameHandler"
	"minesweeper-go/internal/repositories/gameRepository"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.NewAppConfig()
	if err != nil {
		panic(err)
	}

	dbURL := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Container,
		cfg.DB.Name,
	)

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	gameRepository := gameRepository.NewMySQLClient(db)
	gameService := gameService.New(gameRepository)
	gameHandler := gameHandler.NewHTTPHandler(gameService)

	router := gin.New()
	router.GET("/games/:id", gameHandler.Get)
	router.POST("/games", gameHandler.Create)
	router.POST("/games/:id/reveal", gameHandler.Reveal)

	router.Run(":8080")
}
