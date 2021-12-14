package controller

import (
	"net/http"
  "log"

	"github.com/gin-gonic/gin"
  "database/sql"
  _ "github.com/lib/pq"
  "strconv"

	"app/db"
	"app/entity"
)

const (
  queryGetUser    = "SELECT email, name FROM users WHERE id=$1"
  queryGetUserId    = "SELECT id FROM users WHERE email=$1"
  queryInsertUser = "INSERT INTO users (email, name) VALUES($1, $2)"
  queryUpdateUser = "UPDATE users SET email=$1, name=$2 WHERE id=$3"
)

var err error

func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func UpdateUser(c *gin.Context) {
  db := db.Db
  user := entity.User{}

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  logFatal(err)

  c.BindJSON(&user)
  stmt, err := db.Prepare(queryUpdateUser)
  logFatal(err)
  defer stmt.Close()

  _, err = stmt.Exec(user.Email, user.Name, idInt)
  logFatal(err)
  switch {
  default:
    c.JSON(http.StatusOK, gin.H{ "message": id + " has been updated", })
  }
}


func CreateUser(c *gin.Context) {
  db := db.Db
  user := entity.User{}

  c.BindJSON(&user)

  stmt, err := db.Prepare(queryInsertUser)
  logFatal(err)
  defer stmt.Close()

  _, err = stmt.Exec(user.Name, user.Email)
  logFatal(err)

  message := "Create " + user.Name
  c.JSON(http.StatusOK, gin.H{ "message": message, })
}

func GetUser(c *gin.Context){
  db := db.Db
  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  logFatal(err)

  stmt, err := db.Prepare(queryGetUser)
  defer stmt.Close()
  logFatal(err)

  user := entity.User{}
  err = stmt.QueryRow(idInt).Scan(&user.Email, &user.Name)

  // error handling
  switch {
  case err == sql.ErrNoRows:
    c.JSON(http.StatusBadRequest, gin.H{"message": "id " + id + "is not found"})
  case err != nil:
    log.Fatal(err)
  default:
    c.JSON(http.StatusOK, gin.H{"email": user.Email, "name": user.Name})
  }
}

