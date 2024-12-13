package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *Admin) addProfession(c *gin.Context) {
	var input models.Profession
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.ProfessionService.AddProfession(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "",
		"status":     "ok",
		"profession": res,
	})
}

func (h *Admin) updateProfession(c *gin.Context) {
	var input models.Profession
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = id

	err = h.services.ProfessionService.UpdateProfession(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"status":  "ok",
	})
}

func (h *Admin) deleteProfession(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.ProfessionService.DeleteProfession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"status":  "ok",
	})
}

func (h *Admin) getProfessionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res models.Profession

	res, err = h.services.ProfessionService.GetProfessionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    "",
		"status":     "ok",
		"profession": res,
	})
}

func (h *Admin) getProfessions(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	name := c.Query("name")
	code := c.Query("code")
	departmentID, _ := strconv.Atoi(c.Query("department_id"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	res, err := h.services.ProfessionService.GetProfessions(models.ProfessionSearch{
		ID:           id,
		Name:         name,
		Code:         code,
		Limit:        limit,
		DepartmentID: departmentID,
		Page:         page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "",
		"status":      "ok",
		"professions": res,
	})
}
