package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *Admin) addLesson(c *gin.Context) {
	var lesson models.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.LessonService.AddLesson(lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  "ok",
		"lesson":  res,
	})
}

func (h *Admin) updateLesson(c *gin.Context) {
	var lesson models.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("lid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson.ID = id

	err = h.services.LessonService.UpdateLesson(lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

func (h *Admin) deleteLesson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("lid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.LessonService.DeleteLesson(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

func (h *Admin) getLessonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("lid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.LessonService.GetLessonByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "lesson": res})
}

func (h *Admin) getLessons(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	name := c.Query("name")
	code := c.Query("code")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	res, err := h.services.LessonService.GetLessons(models.LessonSearch{
		ID:    id,
		Name:  name,
		Code:  code,
		Limit: limit,
		Page:  page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "lessons": res})
}
