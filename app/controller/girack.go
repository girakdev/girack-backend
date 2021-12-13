package controller

import (
	"net/http"
  "log"

	"github.com/gin-gonic/gin"

	"app/db"
	"app/entity"
)

const (
  queryGetUser    = "SELECT email, name FROM users WHERE id=$1"
)

func DeleteUser(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{ "message": "CreateUser", })
  //Todo
}
func CreateUser(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{ "message": "CreateUser", })
  // Todo
}

func GetUser(c *gin.Context){
  Db := db.GetDB()
  user := entity.User{}

  stmt, err := Db.Prepare(queryGetUser)
  defer stmt.Close()
  logFatal(err)

  err = stmt.QueryRow(1).Scan(&user.Email, &user.Name)
  logFatal(err)

  c.JSON(http.StatusOK, gin.H{"email": user.Email, "name": user.Name})
}

func Index(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{ "message": "Index", })
  // Todo
}
func Show(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{ "message": "Show", })
  // Todo
}
func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}