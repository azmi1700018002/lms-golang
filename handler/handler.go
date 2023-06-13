package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"lms/config/helper"
	"lms/model"
	rauth "lms/repository/r_auth"
)

var identityKey = "id"

func LoginHandler(c *gin.Context) {
	var loginVals model.User
	if err := c.ShouldBind(&loginVals); err != nil {
		c.JSON(400, gin.H{"error": "missing login values"})
		return
	}
	email := loginVals.Email
	password := loginVals.Password

	// Check if user exist
	user, err := rauth.QAuthUser(email, password)

	// Check if username and password match using CheckPasswordHash
	if !helper.CheckPasswordHash(password, user.Password) {
		c.JSON(401, gin.H{"error": "failed authentication"})
		return
	}

	if err != nil {
		c.JSON(401, gin.H{"error": "failed authentication"})
		return
	}

	// Update last login time if authentication succeeds
	if err := rauth.UpdateLastLogin(&user); err != nil {
		c.JSON(401, gin.H{"error": "failed authentication"})
		return
	}

	expireTime := time.Now().Add(1 * time.Hour)

	// konversi Unix ke Time
	expireTimeUTC := time.Unix(expireTime.Unix(), 0).UTC()

	// format waktu sebagai string
	expireTimeString := expireTimeUTC.Format(time.RFC3339)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		identityKey: user.Username,
		"IDUser":    user.IDUser.String(),
		"exp":       expireTime.Unix(),
	}).SignedString([]byte("secret key"))

	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token":   token,
		"IDUser":  user.IDUser.String(),
		"expired": expireTimeString,
		"status":  200,
	})
}
