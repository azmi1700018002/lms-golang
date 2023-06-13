package register

import (
	"net/http"

	"lms/config/db"
	"lms/model"
	rregister "lms/repository/r_register"

	"github.com/gin-gonic/gin"
)

// Make register controller
func Register(g *gin.Context) {
	var user model.User
	var err error

	// Get user data from request body
	if err = g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exist
	if err := db.Server().Where("username = ?", user.Username).First(&user).Error; err == nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "User already exist"})
		return
	}

	// Create new user
	if err = rregister.QRegisterUser(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return Json response success
	g.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})

}
