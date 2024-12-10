package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student-management/internal/models"
)

func (h *V1) signIn(c *gin.Context) {
	var input models.LoginCredentials
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, roleID, err := h.services.AuthService.SignIn(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"status":  "ok",
		"token":   token,
		"role":    roleID,
	})
}
