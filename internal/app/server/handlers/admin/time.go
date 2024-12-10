package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *Admin) addTime(c *gin.Context) {
	var time models.Time
	if err := c.ShouldBindJSON(&time); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.TimeService.AddTime(time)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "time": res})
}

func (h *Admin) updateTime(c *gin.Context) {
	var time models.Time
	if err := c.ShouldBindJSON(&time); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	time.ID = id

	err = h.services.TimeService.UpdateTime(time)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

func (h *Admin) deleteTime(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.TimeService.DeleteTime(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok"})
}

func (h *Admin) getTimeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.TimeService.GetTimeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "time": res})
}

func (h *Admin) getTimes(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	res, err := h.services.TimeService.GetTimes(models.TimeSearch{
		ID:    id,
		Limit: limit,
		Page:  page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "", "status": "ok", "times": res})
}
