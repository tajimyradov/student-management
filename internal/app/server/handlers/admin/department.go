package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *Admin) addDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.DepartmentService.AddDepartment(department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "",
		"status":     "ok",
		"department": res,
	})
}

func (h *Admin) updateDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("did"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department.ID = id

	err = h.services.DepartmentService.UpdateDepartment(department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"status":  "ok",
	})
}

func (h *Admin) deleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("did"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.DepartmentService.DeleteDepartment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"status":  "ok",
	})
}

func (h *Admin) getDepartmentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("did"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var department models.Department

	department, err = h.services.DepartmentService.GetDepartmentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    "",
		"status":     "ok",
		"department": department,
	})
}

func (h *Admin) getDepartments(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	name := c.Query("name")
	code := c.Query("code")
	facultyID, _ := strconv.Atoi(c.DefaultQuery("faculty_id", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	res, err := h.services.DepartmentService.GetDepartments(models.DepartmentSearch{
		ID:        id,
		Name:      name,
		Code:      code,
		FacultyID: facultyID,
		Limit:     limit,
		Page:      page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "",
		"status":      "ok",
		"departments": res,
	})
}

func (h *Admin) uploadFileOfDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("did"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := c.PostForm("name")

	err = h.services.DepartmentService.UploadFile(id, file, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

func (h *Admin) deleteFileOfDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("did"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.DepartmentService.DeleteFile(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

func (h *Admin) getDepartmentInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("did"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.DepartmentService.GetDepartmentInfo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"status":  "ok",
		"info":    res,
	})
}
