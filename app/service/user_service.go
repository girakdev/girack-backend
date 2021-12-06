package service

import (
	"github.com/gin-gin-gonic/gin"

	"app/db"
	"app/entity"
)

type UserService struct{}

func (s UserService) GetAll() (u []entity.User, err error){
  db := db.GetDB()

  u, err = db.Find(&u)

  return
}

func (s UserService) Create(c *gin.Context) (u entity.User, err error){
  db := db.GetDB()
  if err = c.BindJSON(&u); err != nil {
    return
  }

  if err = db.Create(&u); err != nil{
    return
  }

  return
}

func (s UserService) DeleteByEmail(email string) (err error) {
  db := db.GetDB()
  if err = db.Where("id = ?", email).Delete(&entity.User).Error; err != nil {
    return
  }
  return nil
}
