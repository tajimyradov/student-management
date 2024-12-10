package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *Admin) addTimetable(c *gin.Context) {
	var timetable models.Timetable
	if err := c.BindJSON(&timetable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.services.TimetableService.AddTimetable(timetable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "",
	})
}

func (h *Admin) deleteTimetable(c *gin.Context) {
	timetableID, err := strconv.Atoi(c.Param("timetableID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.services.TimetableService.DeleteTimetable(timetableID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
	})
}

func (h *Admin) getTimetableOfGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("group_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	timetables, err := h.services.TimetableService.GetTimetableOfGroup(groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     "ok",
		"message":    "",
		"timetables": timetables,
	})
}
