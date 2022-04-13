package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
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
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": ""})
    return
  }
  defer stmt.Close()
  user.Password, err = passwordHash(user.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": "inviled password format"})
    return
  }

  _, err = stmt.Exec(user.Email, user.Name, user.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": "inviled password format"})
    return
  }
  message := "Create " + user.Name
  c.JSON(http.StatusOK, gin.H{"message": message})
}

func Login(c *gin.Context) {
  session := sessions.Default(c)
  user := entity.User{}
  c.BindJSON(&user)

  dbuser, err := GetUserByEmail(user.Email)

  if err == sql.ErrNoRows {
    c.JSON(http.StatusUnauthorized, gin.H{"error1": "incorrect Email or Password"})
    return
  } else if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "ServerError"})
    return
  }

  err = bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(user.Password))
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect Email or Password"})
    return
  }

  session.Set("userid", user.Id)
  if err = session.Save(); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"errror": "Failed to save session"})
    return
  }
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
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
    return
  }
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
