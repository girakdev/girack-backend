package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
//  "github.com/gin-contrib/sessions"
  "database/sql"
  _ "github.com/lib/pq"
  "strconv"

	"app/db"
	"app/entity"
)

func DeleteUser(c *gin.Context) {
  db := db.Db
  id := c.Param("id")

  idInt, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter must be integer"})
    return
  }

  stmt, err := db.Prepare("DELETE FROM users WHERE id = $1")
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

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter must be STRING"})
    return
  }

  c.BindJSON(&user)
  stmt, err := db.Prepare("UPDATE users SET email=$1, name=$2 WHERE id=$3")
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

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter must be STRING"})
    return
  }

  stmt, err := db.Prepare("SELECT email, name FROM users WHERE id=$1")
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

  stmt, err := db.Prepare("SELECT email, name, id FROM users")
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
    err = rows.Scan(&user.Email, &user.Name, &user.Id)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err})
      return
    }

    users = append(users, user)
  }

  c.JSON(http.StatusOK, users)
}
