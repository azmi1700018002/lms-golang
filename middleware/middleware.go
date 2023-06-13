package middleware

import (
	"lms/model"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var identityKey = "id"

func AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
					"IDUser":    v.IDUser.String(), // tambahkan IDUser ke payload
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Email:  claims[identityKey].(string),
				IDUser: uuid.MustParse(claims["IDUser"].(string)), // konversi string ke uuid.UUID
			}
		},
	})

	if err != nil {
		return nil, err
	}

	if errInit := authMiddleware.MiddlewareInit(); errInit != nil {
		return nil, errInit
	}

	return authMiddleware, nil
}
