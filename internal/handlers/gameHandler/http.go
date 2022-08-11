package gameHandler

import (
	"minesweeper-go/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	gamesService ports.GamesService
}

func NewHTTPHandler(gamesService ports.GamesService) *HTTPHandler {
	return &HTTPHandler{
		gamesService: gamesService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	game, err := hdl.gamesService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, game)
}
