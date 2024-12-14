package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *V1) getFaculties(c *gin.Context) {
	faculties, err := h.services.StudentService.GetFaculties()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, faculties)
}
