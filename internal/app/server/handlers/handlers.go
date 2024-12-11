package handlers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	a "student-management/internal/app/server/handlers/admin"
	v1 "student-management/internal/app/server/handlers/v1"
	"student-management/internal/config"
	"student-management/internal/middlewares"
	"student-management/internal/services"
	"time"
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
		cors.New(
			cors.Config{
				AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
				AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
				AllowCredentials: false,
				AllowAllOrigins:  true,
				MaxAge:           12 * time.Hour,
			},
		),
	)

	//check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1Group := router.Group("/api/v1", middlewares.AuthorizationMiddleware)
	{
		v1Handlers := v1.NewHandler(h.services.ClientService, h.logger, h.config)
		v1Handlers.Init(v1Group)
	}

	adminGroup := router.Group("/admin", middlewares.AuthorizationMiddleware)
	{
		adminHandlers := a.NewHandler(h.services.AdminService, h.logger, h.config)
		adminHandlers.Init(adminGroup)
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
