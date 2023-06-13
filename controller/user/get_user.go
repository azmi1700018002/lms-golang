package user

import (
	"errors"
	ruser "lms/repository/r_user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {
	users, count, err := ruser.MGetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"count": count,
	})
}

func GetUserByID(c *gin.Context) {
	// get the user ID string from the URL parameter
	userIDStr := c.Param("id_user")

	// parse the user ID string into a uuid.UUID value
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// call the MGetUserByID function with the uuid.UUID value
	user, err := ruser.MGetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
