package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *V1) getDepartments(c *gin.Context) {
	facultyId, err := strconv.Atoi(c.DefaultQuery("faculty_id", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	departments, err := h.services.StudentService.GetDepartments(facultyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, departments)
}
