package config

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"lkt_test_interview/cmd/controller"
	"lkt_test_interview/cmd/models"
	"lkt_test_interview/cmd/repository"
	"lkt_test_interview/cmd/services"
	"lkt_test_interview/docs"
)

func StartApp() {
	docs.SwaggerInfo.BasePath = "/"

	db, sqlDB, err := OpenDB()
	if err != nil {
		log.Fatalf("db open: %v", err)
	}

	if err := db.AutoMigrate(&models.Event{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	evRepo := repository.NewEventRepository(db)
	evSvc := services.NewEventService(evRepo)
	evCtrl := controller.NewEventController(evSvc)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	MapUrls(r, evCtrl, sqlDB)

	port := GetEnv("PORT", "8080")
	if port[0] != ':' {
		port = ":" + port
	}

	srv := &http.Server{
		Addr:              port,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on %s", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
