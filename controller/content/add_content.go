package content

import (
	"lms/model"
	"lms/service/s_content"
	"lms/validation/v_content"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentController struct {
	contentService *s_content.ContentService
}

func NewContentController(contentService *s_content.ContentService) *ContentController {
	return &ContentController{contentService: contentService}
}

func (ctrl *ContentController) AddContent(c *gin.Context) {
	var input model.Content
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := v_content.ValidateContent(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err := ctrl.contentService.AddContent(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": content})
}
