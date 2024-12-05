package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *V1) addAuditory(c *gin.Context) {
	var auditory models.Auditory
	if err := c.ShouldBindJSON(&auditory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.AuditoryService.AddAuditory(auditory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"status":   "ok",
		"auditory": res,
	})
}

func (h *V1) updateAuditory(c *gin.Context) {
	var auditory models.Auditory
	if err := c.ShouldBindJSON(&auditory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auditory.ID = id

	err = h.services.AuditoryService.UpdateAuditory(auditory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok"})
}

func (h *V1) deleteAuditory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.AuditoryService.DeleteAuditory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok"})
}

func (h *V1) getAuditoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.AuditoryService.GetAuditoryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok", "auditory": res})
}

func (h *V1) getAuditories(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	name := c.Query("name")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	res, err := h.services.AuditoryService.GetAuditories(models.AuditorySearch{
		ID:    id,
		Name:  name,
		Limit: limit,
		Page:  page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": "ok", "auditories": res})
}
