package admin

import (
	"github.com/gin-gonic/gin"
	"image"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

// AddStudent handler
func (h *Admin) addStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.StudentService.AddStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "student": res})
}

// UpdateStudent handler
func (h *Admin) updateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("sid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student.ID = id
	err = h.services.StudentService.UpdateStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

// DeleteStudent handler
func (h *Admin) deleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.StudentService.DeleteStudent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

// GetStudentByID handler
func (h *Admin) getStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := h.services.StudentService.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "student": student})
}

// GetStudents handler
func (h *Admin) getStudents(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	name := c.Query("name")
	username := c.Query("username")
	code := c.Query("code")
	facultyID, _ := strconv.Atoi(c.DefaultQuery("faculty_id", "0"))
	professionID, _ := strconv.Atoi(c.DefaultQuery("profession_id", "0"))
	groupID, _ := strconv.Atoi(c.DefaultQuery("group_id", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	res, err := h.services.StudentService.GetStudents(models.StudentSearch{
		ID:           id,
		Name:         name,
		FacultyID:    facultyID,
		ProfessionID: professionID,
		GroupID:      groupID,
		Username:     username,
		Code:         code,
		Limit:        limit,
		Page:         page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "students": res})
}

// UploadStudentImage handler
func (h *Admin) uploadStudentImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sid"))
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

	err = h.services.StudentService.UploadImageOfStudent(img, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}
