package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/girakdev/girack-backend/docs"
)

// PingExample godoc
//
//	@Summary	ping
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	ok
//	@Router			/ping [get]
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
