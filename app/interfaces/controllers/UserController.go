package controllers

import (
	"app/domain"
	"app/interfaces/database"
	"app/usecase"
	"strconv"
)

type UserController struct {
  Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
  return &UserController{
    Interactor: usecase.UserInteractor{
      UseRepository: &database.UserRepository {
        SqlHandler: sqlHandler,
      },
    },
  }
}

func (controller *UserController) Create(c Context) {
  u := domain.User{}
  c.Bind(&u)
  err := controller.interactor.Add(u)
  if err != nil {
    c.JSON(500, NewError(err))
    return
  }
  c.Json(201)
}

func (controller *UserController) Show(c Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  user, err := controller.Interactor.UsreById(id)
  if err != nil {
    c.JSON(500, NewError(err))
  }
  c.JSON(200, user)
}
