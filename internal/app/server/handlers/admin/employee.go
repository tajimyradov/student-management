package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"student-management/internal/models"
)

func (h *Admin) addEmployeeRate(c *gin.Context) {
	var input models.EmployeeRate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.services.EmployeeService.AddEmployeeRate(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
	})
}

func (h *Admin) deleteEmployeeRate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("emp_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.EmployeeService.DeleteEmployeeRate(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
	})
}

func (h *Admin) updateEmployeeRate(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("emp_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.EmployeeRate
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = id

	err = h.services.EmployeeService.UpdateEmployeeRate(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
	})
}

func (h *Admin) getEmployeeRate(c *gin.Context) {
	res, err := h.services.EmployeeService.GetAllEmployeeRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "",
		"status":        "ok",
		"employee_rate": res,
	})
}

func (h *Admin) getPositions(c *gin.Context) {
	res, err := h.services.EmployeeService.GetAllPositions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "",
		"status":    "ok",
		"positions": res,
	})
}

func (h *Admin) getEmployeeRateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("emp_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.services.EmployeeService.GetEmployeeRatByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"employee_rate": res,
		"status":        "ok",
		"message":       "",
	})
}
