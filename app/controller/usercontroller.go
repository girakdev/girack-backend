package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
  "database/sql"
  _ "github.com/lib/pq"
  "strconv"

	"app/db"
	"app/entity"
)

func DeleteUser(c *gin.Context) {
  db := db.Db
  id := c.Param("id")
  query := "DELETE FROM users WHERE id = $1"

  idInt, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter must be integer"})
    return
  }

  stmt, err := db.Prepare(query)
  defer stmt.Close()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make Statement"})
    return
  }

  _, err = stmt.Exec(idInt)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    return
  }
  c.JSON(http.StatusOK, gin.H{ "message": id + " has been deleted",})
}

func UpdateUser(c *gin.Context) {
  db := db.Db
  user := entity.User{}
  query := "UPDATE users SET email=$1, name=$2 WHERE id=$3"

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter must be STRING"})
    return
  }

  c.BindJSON(&user)
  stmt, err := db.Prepare(query)
  defer stmt.Close()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    return
  }

  _, err = stmt.Exec(user.Email, user.Name, idInt)

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    return
  }
  c.JSON(http.StatusOK, gin.H{ "message": id + " has been updated"})
}

func GetUser(c *gin.Context){
  db := db.Db
  query := "SELECT email, name FROM users WHERE id=$1"

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter must be STRING"})
    return
  }


  stmt, err := db.Prepare(query)
  defer stmt.Close()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    return
  }

  user := entity.User{}
  err = stmt.QueryRow(idInt).Scan(&user.Email, &user.Name)

  if err == sql.ErrNoRows {
    c.JSON(http.StatusNotFound, gin.H{"error": "id " + id + " is not found"})
    return
  } else if err != nil{
    c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    return
  }
  c.JSON(http.StatusOK, user)
}

func GetAllUser(c *gin.Context) {
  db := db.Db
  query := "SELECT email, name FROM users"

  stmt, err := db.Prepare(query)
  defer stmt.Close()

  rows, err := stmt.Query()
  if err == sql.ErrNoRows {
    c.JSON(http.StatusNotFound, gin.H{"error": "we have no users ;;",})
    return
  } else if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    return
  }

  users := []entity.User{}
  for rows.Next() {
    user := entity.User{}
    err = rows.Scan(&user.Email, &user.Name)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err})
      return
    }

    users = append(users, user)
  }

  c.JSON(http.StatusOK, users)
}
