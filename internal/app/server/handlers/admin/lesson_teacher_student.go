package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *Admin) addLessonTeacherStudentBinding(c *gin.Context) {
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

func (h *Admin) deleteLessonTeacherStudentBinding(c *gin.Context) {
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

func (h *Admin) getLessonTeacherStudentBinding(c *gin.Context) {
	teacherId, err := strconv.Atoi(c.Query("teacherId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lessonId, err := strconv.Atoi(c.Query("lessonId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
