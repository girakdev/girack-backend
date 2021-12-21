package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gin-contrib/sessions"
	"golang.org/x/crypto/bcrypt"
  "strings"

	"app/db"
)

func Login(c *gin.Context) {
  session := sessions.Default(c)
  email := c.PostForm("email")
  password := c.PostForm("password")
  if strings.Trim(email, " ") == "" || strings.Trim(password, " ") == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
    return
  }

  hashpassword, _ := passwordHash(password)
  dbpassword, err := GetPasswordByEmail(email)
  if err != sql.ErrNoRows {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect Email or Password"})
    return
  } else if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "EEEEEEEEEEEEEEEERROR"})
    return
  }

  err = passwordVerify(hashpassword, dbpassword)
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect Email or Password"})
    return
  }

  session.Set("user", email)
  if err = session.Save(); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"errror": "Failed to save session"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func Logout(c *gin.Context) {
  session := sessions.Default(c)
  user := session.Get("user")
  if user == nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session toekn"})
  }
  session.Clear()
  if err := session.Save(); err != nil {
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

func passwordVerify(hash, pw string) error {
  return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func GetPasswordByEmail(email string) (password string, err error) {
  db := db.Db
  query := "SELECT password FROM users WHERE email = $1"

  stmt, err := db.Prepare(query)
  if err != nil {
    return "", err
  }
  err = stmt.QueryRow(email).Scan(&password)
  if err != nil {
    return "", err
  }
  return
}

