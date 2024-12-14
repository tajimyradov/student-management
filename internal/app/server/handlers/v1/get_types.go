package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *V1) getTypes(c *gin.Context) {
	types, err := h.services.StudentService.GetTypes()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, types)
}
