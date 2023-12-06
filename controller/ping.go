package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/girakdev/girack-backend/docs"
)

// @BasePath /v1

// PingExample godoc
// @Summary ping
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")

}
