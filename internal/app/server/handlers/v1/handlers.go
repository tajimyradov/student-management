package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"student-management/internal/config"
	"student-management/internal/services/client"
)

type V1 struct {
	services *client.Service
	logger   *zap.Logger
	config   *config.AppConfig
}

func NewHandler(services *client.Service, logger *zap.Logger, config *config.AppConfig) *V1 {
	return &V1{
		services: services,
		logger:   logger,
		config:   config,
	}
}

func (h *V1) Init(v1 *gin.RouterGroup) {
	v1.POST("/sign-in", h.signIn)
	v1.GET("/faculties", h.getFaculties)
	v1.GET("/departments", h.getDepartments)
	v1.GET("/groups", h.getGroups)
	v1.GET("/students", h.getStudents)
	v1.GET("/lessons", h.getLessons)
	v1.GET("/types", h.getTypes)
	v1.GET("/times", h.getTimes)
	v1.POST("/check-in", h.checkStudentsForExistence)
}
