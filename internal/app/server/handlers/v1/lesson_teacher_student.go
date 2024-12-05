package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student-management/internal/models"
)

func (h *V1) addLessonTeacherStudentBinding(c *gin.Context) {
	var input models.LessonTeacherStudent
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.services.TimetableService.AddStudentTeacherLessonBinding(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
	})
}

func (h *V1) deleteLessonTeacherStudentBinding(c *gin.Context) {
	var input models.LessonTeacherStudent
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.services.TimetableService.DeleteStudentTeacherLessonBinding(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
	})
}
