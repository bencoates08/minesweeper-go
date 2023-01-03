package gameHandler

import (
	"log"
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
	game, err := hdl.gamesService.Get(c, c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, BuildGameResponse(game))
}

func (hdl *HTTPHandler) Create(c *gin.Context) {
	var request CreateRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(400, gin.H{"message": "Error binding json body: " + err.Error()})
		return
	}
	game, err := hdl.gamesService.Create(c, request.Name, request.Height, request.Width, request.Bombs)
	if err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(500, gin.H{"message": "Error creating game: " + err.Error()})
		return
	}

	c.JSON(201, BuildGameResponse(game))

	log.Printf("Game created: %s", game.ID)
}

func (hdl *HTTPHandler) Reveal(c *gin.Context) {
	var request RevealRequest
	c.BindJSON(&request)
	game, err := hdl.gamesService.Reveal(c, c.Param("id"), request.Row, request.Col)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, BuildGameResponse(game))
}
