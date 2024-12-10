package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *V1) getStudents(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lessonID, err := strconv.Atoi(c.Query("lesson_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	typeID, err := strconv.Atoi(c.Query("type_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	claims := value.(models.UserClaims)

	students, err := h.services.StudentService.GetStudents(claims.UserID, claims.RoleID, lessonID, groupID, typeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "",
		"status":   "ok",
		"students": students,
	})
}
