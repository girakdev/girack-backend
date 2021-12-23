package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/sessions"
  "database/sql"
  _ "github.com/lib/pq"

	"app/db"
	"app/entity"
)

func CreateChannel(c *gin.Context) {
  db := db.Db
  session := sessions.Default(c)
  channel := entity.Channel{}
  c.BindJSON(&channel)

  stmt, err := db.Prepare("INSERT INTO channels (name, description, dm_flag, member) VALUES ($1, $2, $3, '{$4}')")
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err})
  }
  defer stmt.Close()

  creater := session.Get("userid")
  if creater == nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session Token"})
    return
  }
  _, err = stmt.Exec(channel.Name, channel.Description, channel.Dm_flag, creater)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Created Channel"})
}

func UpdateChannel(c *gin.Context) {
}

func GetChannel(c *gin.Context) {
}

func GetAllChannel(c *gin.Context) {
}

func DeleteChannel(c *gin.Context) {
}

