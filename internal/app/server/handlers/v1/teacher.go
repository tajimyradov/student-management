package v1

import (
	"github.com/gin-gonic/gin"
	"image"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

// AddTeacher handler
func (h *V1) addTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.TeacherService.AddTeacher(teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok", "teacher": res})
}

// UpdateTeacher handler
func (h *V1) updateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher.ID = id
	err = h.services.TeacherService.UpdateTeacher(teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok"})
}

// DeleteTeacher handler
func (h *V1) deleteTeacher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.TeacherService.DeleteTeacher(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok"})
}

// GetTeacherByID handler
func (h *V1) getTeacherByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher, err := h.services.TeacherService.GetTeacherByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok", "teacher": teacher})
}

// GetTeachers handler
func (h *V1) getTeachers(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")
	username := c.Query("username")
	code := c.Query("code")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	res, err := h.services.TeacherService.GetTeachers(models.TeacherSearch{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Code:      code,
		Limit:     limit,
		Page:      page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok", "teachers": res})
}

// UploadTeacherImage handler
func (h *V1) uploadTeacherImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid image format"})
		return
	}

	err = h.services.TeacherService.UploadImageOfTeacher(img, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok"})
}
