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
	v1.GET("/timetable", h.getTimetable)
	v1.GET("/students", h.getStudents)
	v1.POST("/check-in", h.checkStudentsForExistence)
}
