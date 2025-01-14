package admin

import (
	"github.com/gin-gonic/gin"
	"image"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

// AddTeacher handler
func (h *Admin) addTeacher(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "teacher": res})
}

// UpdateTeacher handler
func (h *Admin) updateTeacher(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

// DeleteTeacher handler
func (h *Admin) deleteTeacher(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

// GetTeacherByID handler
func (h *Admin) getTeacherByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "teacher": teacher})
}

// GetTeachers handler
func (h *Admin) getTeachers(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	name := c.Query("name")
	username := c.Query("username")
	code := c.Query("code")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	departmentID, _ := strconv.Atoi(c.DefaultQuery("department_id", "0"))
	facultyID, _ := strconv.Atoi(c.DefaultQuery("faculty_id", "0"))

	res, err := h.services.TeacherService.GetTeachers(models.TeacherSearch{
		ID:           id,
		FacultyID:    facultyID,
		Name:         name,
		Username:     username,
		Code:         code,
		DepartmentId: departmentID,
		Limit:        limit,
		Page:         page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "teachers": res})
}

// UploadTeacherImage handler
func (h *Admin) uploadTeacherImage(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}
