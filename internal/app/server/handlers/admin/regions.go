package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Admin) getRegions(c *gin.Context) {
	res, err := h.services.RegionService.GetRegions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
		"regions": res,
	})
}
