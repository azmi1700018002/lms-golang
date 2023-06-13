package content

import (
	"lms/service/s_content"
	"lms/validation/v_content"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func DeleteContent(c *gin.Context) {
// 	idcontent, err := strconv.Atoi(c.Param("idcontent"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Content ID"})
// 		return
// 	}

// 	err = rcontent.MDeleteContentByID(idcontent)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": "Content deleted"})
// }

type contentDeleteController struct {
	contentDeleteService s_content.ContentDeleteService
}

func DeleteContentController(contentDeleteService s_content.ContentDeleteService) *contentDeleteController {
	return &contentDeleteController{contentDeleteService}
}

func (c *contentDeleteController) DeleteContent(ctx *gin.Context) {
	idcontent := ctx.Param("idcontent")
	id, err := v_content.ValidateDeleteID(idcontent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.contentDeleteService.DeleteContent(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Content deleted"})
}
