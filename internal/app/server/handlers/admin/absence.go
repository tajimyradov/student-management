package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *Admin) getAbsences(c *gin.Context) {
	var absence models.AbsenceSearch
	if err := c.ShouldBindJSON(&absence); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.TimetableService.GetAbsences(absence)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Admin) updateAbsences(c *gin.Context) {
	var absence models.Absence
	if err := c.ShouldBindJSON(&absence); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.services.TimetableService.UpdateAbsence(absence.Status, absence.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, absence)
}

func (h *Admin) SyncAbsence(c *gin.Context) {
	//err := h.services.TimetableService.Sync()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
	})
}

func (h *Admin) getAbsenceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("absence_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.TimetableService.GetAbsenceByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
