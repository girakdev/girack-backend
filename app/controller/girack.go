package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"app/db"
	"app/entity"
)

const (
  queryGetUser    = "SELECT * FROM users WHERE id=$1"
)

func DeleteUser(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{ "message": "CreateUser", })
  //Todo
}
func CreateUser(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{ "message": "CreateUser", })
  // Todo
}

func GetUser(){
  Db := db.GetDB()

  rows, err := Db.Query("SELECT * FROM users")
  checkErr(err)
  defer rows.Close()

  var user entity.User

  for rows.Next() {
    err := rows.Scan(&user.Email, &user.Name)
    checkErr(err)
    fmt.Printf("Email: %s, Name: %s\n", user.Email, user.Name)
  }
  err = rows.Err()
  checkErr(err)

  //err = stmt.QueryRow(1).Scan(&user.Email, &user.Name)
  //checkErr(err)

//  c.JSON(http.StatusOK, gin.H{"email": user.Email, "name": user.Name})
}

func Index(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{ "message": "Index", })
  // Todo
}
func Show(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{ "message": "Show", })
  // Todo
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
