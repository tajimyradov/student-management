package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *V1) getTimes(c *gin.Context) {
	times, err := h.services.StudentService.GetTimes()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, times)
}
