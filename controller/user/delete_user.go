package user

import (
	"lms/service/s_user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userDeleteController struct {
	userDeleteService s_user.UserDeleteService
}

func NewUserDeleteController(userDeleteService s_user.UserDeleteService) *userDeleteController {
	return &userDeleteController{userDeleteService}
}

func (c *userDeleteController) DeleteUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id_user"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.userDeleteService.DeleteUserByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "User deleted"})
}
