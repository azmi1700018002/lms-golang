package controller

import "github.com/gin-gonic/gin"

func Helloworld(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "It works",
	})
}
