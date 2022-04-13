package controller

import (
  "net/http"

	"github.com/gin-gonic/gin"
//"github.com/gin-contrib/sessions"
//  "database/sql"
  _ "github.com/lib/pq"
  "strconv"

	"app/db"
	"app/entity"
)

const (
  deleteUserQuery = "DELETE FROM users WHERE id = $1"
  updateUserQuery = "UPDATE users SET email=$1, name=$2 WHERE id=$3"
  getUserQuery    = "SELECT email, name FROM users WHERE id=$1"
  getAllUserQuery = "SELECT email, name, id FROM users"
)

func DeleteUser(c *gin.Context) {
  db := db.Db
  id := c.Param("id")

  idInt, err := strconv.Atoi(id)
  CheckError(c, err, http.StatusInternalServerError,)

  stmt, err := db.Prepare(deleteUserQuery)
  defer stmt.Close()
  CheckError(c, err, http.StatusInternalServerError,)

  _, err = stmt.Exec(idInt)
  CheckError(c, err, http.StatusInternalServerError,)

  c.JSON(http.StatusOK, gin.H{ "message": id + " has been deleted",})
}

func UpdateUser(c *gin.Context) {
  db := db.Db
  user := entity.User{}

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)

  CheckError(c, err, http.StatusInternalServerError,)

  c.BindJSON(&user)
  stmt, err := db.Prepare(updateUserQuery)
  defer stmt.Close()
  CheckError(c, err, http.StatusInternalServerError,)

  _, err = stmt.Exec(user.Email, user.Name, idInt)
  CheckError(c, err, http.StatusInternalServerError,)

  c.JSON(http.StatusOK, gin.H{ "message": id + " has been updated"})
}

func GetUser(c *gin.Context){
  db := db.Db

  id := c.Param("id")
  idInt, err := strconv.Atoi(id)

  CheckError(c, err, http.StatusInternalServerError,)

  stmt, err := db.Prepare(getUserQuery)
  defer stmt.Close()
  CheckError(c, err, http.StatusInternalServerError,)

  user := entity.User{}
  err = stmt.QueryRow(idInt).Scan(&user.Email, &user.Name)
  CheckError(c, err, http.StatusInternalServerError,)

  c.JSON(http.StatusOK, user)
}

func GetAllUser(c *gin.Context) {
  db := db.Db

  stmt, err := db.Prepare(getAllUserQuery)
  defer stmt.Close()
  CheckError(c, err, http.StatusInternalServerError,)

  rows, err := stmt.Query()
  CheckError(c, err, http.StatusInternalServerError,)

  users := []entity.User{}

  for rows.Next() {
    user := entity.User{}

    err = rows.Scan(&user.Email, &user.Name, &user.Id)
    CheckError(c, err, http.StatusInternalServerError,)
    users = append(users, user)
  }
  c.JSON(http.StatusOK, users)
}
