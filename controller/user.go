package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/girakdev/girack-backend/application/model"
	"github.com/girakdev/girack-backend/application/usecase"
	"github.com/girakdev/girack-backend/internal/pulid"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

// @Summary	Get User
// @Schemes
// @Description	Get User
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	model.User
// @Router			/users [get]
func (c *UserController) GetUser(ctx *gin.Context) {
	guOut, err := c.userUsecase.GetUser(ctx, &usecase.GetUserInput{
		ID: pulid.ID(ctx.Param("id")),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user := model.User{
		ID:   guOut.User.ID,
		Name: guOut.User.Name,
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary	List User
// @Schemes
// @Description	Get List User
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]model.User
// @Router			/users [get]
func (c *UserController) ListUser(ctx *gin.Context) {
	// TODO: Implement

	ctx.JSON(http.StatusOK, gin.H{"message": "List User"})
}

// @Summary	Create User
// @Schemes
// @Description	Create User
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	model.User
// @Router			/users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	// TODO: Implement

	ctx.JSON(http.StatusOK, gin.H{"message": "Create User"})
}

// @Summary	Delete User
// @Schemes
// @Description	Delete User
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	model.User
// @Router			/users [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	// TODO: Implement

	ctx.JSON(http.StatusOK, gin.H{"message": "Delete User"})
}

// @Summary	Update User
// @Schemes
// @Description	Update User
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	model.User
// @Router			/users [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	// TODO: Implement

	ctx.JSON(http.StatusOK, gin.H{"message": "Update User"})
}
