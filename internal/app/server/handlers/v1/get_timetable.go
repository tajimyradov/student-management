package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student-management/internal/models"
	"time"
)

func (h *V1) getTimetable(c *gin.Context) {
	weekday := time.Now().Weekday().String()
	value, ok := c.Get("claims")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	claims := value.(models.UserClaims)
	timetable, err := h.services.TimetableService.GetTimetable(claims.UserID, claims.RoleID, weekday)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "",
		"status":    "ok",
		"timetable": timetable,
	})
}
