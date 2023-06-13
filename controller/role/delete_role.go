package role

import (
	rrole "lms/repository/r_role"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRole(c *gin.Context) {
	id_role, err := strconv.Atoi(c.Param("id_role"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Role ID"})
		return
	}

	err = rrole.MDeleteRolesByID(id_role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Role deleted"})
}
