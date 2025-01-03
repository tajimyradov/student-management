package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Admin) getStatisticsByGender(c *gin.Context) {
	facultyID, err := strconv.Atoi(c.Param("faculty_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.StatisticsService.GetStatisticsByGender(facultyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Admin) getStatisticsOfProfession(c *gin.Context) {
	facultyID, err := strconv.Atoi(c.Param("faculty_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.StatisticsService.GetStatisticsProfession(facultyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Admin) getStatisticsByAge(c *gin.Context) {
	facultyID, err := strconv.Atoi(c.Param("faculty_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.StatisticsService.GetStatisticsByAge(facultyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Admin) getStatisticsByRegions(c *gin.Context) {
	facultyID, err := strconv.Atoi(c.Param("faculty_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.StatisticsService.GetStatisticsByRegions(facultyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
