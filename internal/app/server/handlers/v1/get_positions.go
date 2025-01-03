package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *V1) getPositions(c *gin.Context) {
	res, err := h.services.StudentService.GetPositions()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
