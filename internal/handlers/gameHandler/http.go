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

	c.JSON(200, BuildGameResponse(game))
}

func (hdl *HTTPHandler) Create(c *gin.Context) {
	var request CreateRequest
	c.BindJSON(&request)
	game, err := hdl.gamesService.Create(request.Name, request.Height, request.Width, request.Bombs)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, BuildGameResponse(game))
}

func (hdl *HTTPHandler) Reveal(c *gin.Context) {
	var request RevealRequest
	c.BindJSON(&request)
	game, err := hdl.gamesService.Reveal(c.Param("id"), request.Row, request.Col)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, BuildGameResponse(game))
}
