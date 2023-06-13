package role

import (
	"lms/model"
	rrole "lms/repository/r_role"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRole(c *gin.Context) {
	var input model.Role
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id_role, err := strconv.Atoi(c.Param("id_role"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Role ID"})
		return
	}

	_, err = rrole.MUpdateRolesByID(id_role, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role updated"})
}
