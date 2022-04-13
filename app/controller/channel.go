package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/sessions"
//  "database/sql"
  "strconv"
  _ "github.com/lib/pq"

	"app/db"
	"app/entity"
)

const (
  createChannelQuery = "INSERT INTO channels (name, description, dm_flag, member) VALUES ($1, $2, $3, '{$4}')"
  updateChannelQuery = "UPDATE channels SET name=$1, description=$2, member=$3 WHERE id=$4"
)

func CreateChannel(c *gin.Context) {
  db := db.Db
  session := sessions.Default(c)
  channel := entity.Channel{}
  c.BindJSON(&channel)

  stmt, err := db.Prepare(createChannelQuery)
  defer stmt.Close()
  CheckError(c, err, http.StatusInternalServerError)

  creater := session.Get("userid")
  _, err = stmt.Exec(channel.Name, channel.Description, channel.Dm_flag, creater)
  CheckError(c, err, http.StatusInternalServerError)

  c.JSON(http.StatusOK, gin.H{"message": "Channel Created "})
}


func UpdateChannel(c *gin.Context) {
  db := db.Db
  channel := entity.Channel{}

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  CheckError(c, err, http.StatusInternalServerError)

  c.BindJSON(&channel)
  stmt, err := db.Prepare(updateChannelQuery)
  defer stmt.Close()
  CheckError(c, err, http.StatusInternalServerError)

  _, err = stmt.Exec(channel.Name, channel.Description, channel.Member, idInt)
  CheckError(c, err, http.StatusInternalServerError)

  c.JSON(http.StatusOK, gin.H{ "message": channel.Name + " has been updated"})

}

func GetChannel(c *gin.Context) {
}

func GetAllChannel(c *gin.Context) {
}

func DeleteChannel(c *gin.Context) {
}

