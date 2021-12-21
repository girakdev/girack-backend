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

func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}


func DeleteUser(c *gin.Context) {
  db := db.Db
  id := c.Param("id")
  query := "DELETE FROM users WHERE id = $1"

  idInt, err := strconv.Atoi(id)
  logFatal(err)

  stmt, err := db.Prepare(query)
  logFatal(err)
  defer stmt.Close()

  _, err = stmt.Exec(idInt)
  switch {
  case err != nil:
    panic(err)
  default:
    c.JSON(http.StatusNoContent, gin.H{ "message": id + " has been deleted",})
  }
}

func UpdateUser(c *gin.Context) {
  db := db.Db
  user := entity.User{}
  query := "UPDATE users SET email=$1, name=$2 WHERE id=$3"

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  logFatal(err)

  c.BindJSON(&user)
  stmt, err := db.Prepare(query)
  logFatal(err)
  defer stmt.Close()

  _, err = stmt.Exec(user.Email, user.Name, idInt)

  switch {
  case err != nil:
    log.Fatal(err)
  default:
    c.JSON(http.StatusOK, gin.H{ "message": id + " has been updated", })
  }
}


func CreateUser(c *gin.Context) {
  db := db.Db
  user := entity.User{}
  query := "INSERT INTO users (email, name) VALUES($1, $2)"

  c.BindJSON(&user)

  stmt, err := db.Prepare(query)
  logFatal(err)
  defer stmt.Close()

  _, err = stmt.Exec(user.Name, user.Email)
  logFatal(err)

  message := "Create " + user.Name
  c.JSON(http.StatusCreated, gin.H{ "message": message, })
}

func GetUser(c *gin.Context){
  db := db.Db
  query := "SELECT email, name FROM users WHERE id=$1"

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  logFatal(err)

  stmt, err := db.Prepare(query)
  defer stmt.Close()
  logFatal(err)

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

func GetAllUser(c *gin.Context) {
  db := db.Db
  query := "SELECT email, name FROM users"

  stmt, err := db.Prepare(query)
  logFatal(err)
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
    logFatal(err)

    users = append(users, user)
  }

  c.JSON(http.StatusOK, users)
}
