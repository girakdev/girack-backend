package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/girakdev/girack-backend/application/usecase"
	"github.com/girakdev/girack-backend/infrastructure/datastore"
)

func Serve() {
	// Repository
	channelRepository := datastore.NewChannelRepository()

	// Usecase
	channelUsecase := usecase.NewChannnelUsecase(channelRepository)

	// Handler
	channelController := NewChannelHandler(channelUsecase)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/channels", channelController.ListChannel)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")gin.Bind(j:w)
}
