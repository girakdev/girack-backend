package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
  //"database/sql"
	_ "github.com/lib/pq"
	"github.com/gin-contrib/sessions"
	"golang.org/x/crypto/bcrypt"

	"app/db"
  "app/entity"
)

const (
  registerQuery = "INSERT INTO users (email, name, password) VALUES($1, $2, $3)"
  getUserByEmailQuery = "SELECT id, email, password, name FROM users WHERE email = $1"
)

func Register(c *gin.Context) {
  db := db.Db
  user := entity.User{}
  c.BindJSON(&user)

  stmt, err := db.Prepare(registerQuery)
  defer stmt.Close()
  CheckError(c, err, http.StatusInternalServerError)

  user.Password, err = passwordHash(user.Password)
  CheckError(c, err, http.StatusInternalServerError)

  _, err = stmt.Exec(user.Email, user.Name, user.Password)

  CheckError(c, err, http.StatusInternalServerError)
  message := "Create " + user.Name

  c.JSON(http.StatusOK, gin.H{"message": message})
}

func Login(c *gin.Context) {
  session := sessions.Default(c)
  user := entity.User{}
  c.BindJSON(&user)

  dbuser, err := GetUserByEmail(user.Email)
  CheckError(c, err, http.StatusInternalServerError)

  err = bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(user.Password))
  CheckError(c, err, http.StatusInternalServerError)

  session.Set("userid", user.Id)
  err = session.Save()
  CheckError(c, err, http.StatusInternalServerError)

  c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func Logout(c *gin.Context) {
  session := sessions.Default(c)
  user := session.Get("userid")

  if user == nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session toekn"})
    return
  }

  session.Delete("userid")
  err := session.Save()
  CheckError(c, err, http.StatusInternalServerError)

  c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}


func passwordHash(pw string) (string, error) {
  hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
  if err != nil {
    return "", err
  }
  return string(hash), err
}

func GetUserByEmail(email string) (user entity.User, err error) {
  db := db.Db

  stmt, err := db.Prepare(getUserByEmailQuery)
  if err != nil {
    return
  }
  err = stmt.QueryRow(email).Scan(&user.Id, &user.Email, &user.Password, &user.Name)
  return
}
