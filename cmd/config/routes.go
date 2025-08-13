package config

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"lkt_test_interview/cmd/controller"
)

func MapUrls(r *gin.Engine, ev *controller.EventController, sqlDB interface{ PingContext(context.Context) error }) {
	r.GET("/livez", func(c *gin.Context) { c.Status(http.StatusOK) })
	r.GET("/readyz", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := sqlDB.PingContext(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "db not ready", "details": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api/Events")
	{
		api.POST("", ev.Create)
		api.GET("", ev.List)
		api.GET("/:id", ev.GetByID)
	}

	r.GET("/ping", controller.PingController.Ping)
}
