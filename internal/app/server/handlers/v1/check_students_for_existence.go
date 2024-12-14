package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student-management/internal/models"
)

func (h *V1) checkStudentsForExistence(c *gin.Context) {
	var input models.Absence
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	claims := value.(models.UserClaims)
	input.TeacherID = claims.UserID

	if !(claims.RoleID == 2 || claims.RoleID == 3) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	err := h.services.StudentService.CheckForExistence(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
	})
}
