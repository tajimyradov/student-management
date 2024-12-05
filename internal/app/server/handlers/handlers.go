package handlers

import (
	"fmt"
	"net/http"
	v1 "student-management/internal/app/server/handlers/v1"
	"student-management/internal/config"
	"student-management/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	logger   *zap.Logger
	config   *config.AppConfig
	services *services.Service
}

func NewHandler(
	logger *zap.Logger,
	config *config.AppConfig,
	services *services.Service,
) *Handler {
	return &Handler{
		config:   config,
		logger:   logger,
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	// Init gin handler
	router := gin.Default()
	router.Use(
		RequestCancelRecover(),
		cors.Default(),
	)

	//check endpoint
	router.GET("/ping", func(c *gin.Context) {

		c.String(http.StatusOK, "pong")
	})

	admin := router.Group("/admin")
	{
		v1Handlers := v1.NewHandler(h.services, h.logger, h.config)
		v1Handlers.Init(admin)
	}

	return router
}

func RequestCancelRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("client cancel the request")
				c.Request.Context().Done()
			}
		}()
		c.Next()
	}
}
