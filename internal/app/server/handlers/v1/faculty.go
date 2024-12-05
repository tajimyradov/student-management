package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *V1) addFaculty(c *gin.Context) {
	var faculty models.Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.FacultyService.AddFaculty(faculty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"status":  "ok",
		"faculty": res,
	})
}

func (h *V1) updateFaculty(c *gin.Context) {
	var faculty models.Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("fid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	faculty.ID = id

	err = h.services.FacultyService.UpdateFaculty(faculty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  "ok",
	})
}

func (h *V1) deleteFaculty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("fid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.FacultyService.DeleteFaculty(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  "ok",
	})
}

func (h *V1) getFacultyByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("fid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var faculty models.Faculty

	faculty, err = h.services.FacultyService.GetFacultyByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  "ok",
		"faculty": faculty,
	})
}

func (h *V1) getFaculties(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	name := c.Query("name")
	code := c.Query("code")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	res, err := h.services.FacultyService.GetFaculties(models.FacultySearch{
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

	c.JSON(http.StatusOK, gin.H{
		"message":   "ok",
		"status":    "ok",
		"faculties": res,
	})
}
