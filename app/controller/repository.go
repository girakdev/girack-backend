package controller

import (
  //"net/http"
  "github.com/gin-gonic/gin"
)

func CheckError(c *gin.Context, err error, status int) {
  if err != nil {
    c.JSON(status, gin.H{"message": err.Error()})
    c.Abort()
  }
}
