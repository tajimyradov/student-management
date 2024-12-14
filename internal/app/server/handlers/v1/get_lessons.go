package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *V1) getLessons(c *gin.Context) {
	lessons, err := h.services.StudentService.GetLessons()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, lessons)
}
