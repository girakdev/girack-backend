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
  queryGetChannel    = "SELECT email, name FROM channels WHERE id=$1"
  queryGetAllChannel = "SELECT * FROM users"
  queryInsertChannel = "INSERT INTO users (email, name) VALUES($1, $2)"
  queryUpdateChannel = "UPDATE users SET email=$1, name=$2 WHERE id=$3"
  queryDeleteChannel = "DELETE FROM users WHERE id = $1"
)

func DeleteChannel(c *gin.Context) {
  db := db.Db
  id := c.Param("id")

  idInt, err := strconv.Atoi(id)
  if err != nil {
    log.Fatal(err)
  }

  stmt, err := db.Prepare(queryDeleteChannel)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(idInt)
  switch {
  case err != nil:
    log.Fatal(err)
  default:
    c.JSON(http.StatusNoContent, gin.H{ "message": id + " has been deleted",})
  }
}

func UpdateChannel(c *gin.Context) {
  db := db.Db
  user := entity.User{}

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  if err != nil {
    log.Fatal(err)
  }

  c.BindJSON(&user)
  stmt, err := db.Prepare(queryUpdateChannel)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(user.Email, user.Name, idInt)

  switch {
  case err != nil:
    log.Fatal(err)
  default:
    c.JSON(http.StatusOK, gin.H{ "message": id + " has been updated", })
  }
}


func CreateChannel(c *gin.Context) {
  db := db.Db
  user := entity.User{}

  c.BindJSON(&user)

  stmt, err := db.Prepare(queryInsertChannel)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(user.Name, user.Email)
  if err != nil {
    log.Fatal(err)
  }

  message := "Create " + user.Name
  c.JSON(http.StatusCreated, gin.H{ "message": message, })
}

func GetChannel(c *gin.Context){
  db := db.Db
  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  if err != nil {
    log.Fatal(err)
  }

  stmt, err := db.Prepare(queryGetChannel)
  defer stmt.Close()
  if err != nil {
    log.Fatal(err)
  }

  user := entity.User{}
  err = stmt.QueryRow(idInt).Scan(&user.Email, &user.Name)

  switch {
  case err == sql.ErrNoRows:
    c.JSON(http.StatusNotFound, gin.H{"message": "id " + id + " is not found"})
  case err != nil:
    log.Fatal(err)
  default:
    c.JSON(http.StatusOK, user)
  }
}

func GetAllChannel(c *gin.Context) {
  db := db.Db

  stmt, err := db.Prepare(queryGetAllChannel)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  rows, err := stmt.Query()
  switch {
  case err == sql.ErrNoRows:
    c.JSON(http.StatusNotFound, gin.H{"message": "we have no users ;;",})
  case err != nil:
    log.Fatal(err)
  }

  users := []entity.User{}
  for rows.Next() {
    user := entity.User{}
    err = rows.Scan(&user.Email, &user.Name)
    if err != nil {
    log.Fatal(err)
    }
    users = append(users, user)
  }

  c.JSON(http.StatusOK, users)
}
