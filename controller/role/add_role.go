package role

import (
	"lms/model"
	rrole "lms/repository/r_role"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoles(c *gin.Context) {
	var input model.Role
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roles, err := rrole.MaddRoles(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roles})
}
