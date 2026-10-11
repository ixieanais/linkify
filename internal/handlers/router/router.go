// Package router
package router

import (
	"context"
	"time"

	"linkify/internal/config"
	"linkify/internal/database"
	v1 "linkify/internal/handlers/v1"
	"linkify/internal/middlewares"
	"linkify/internal/repositories"
	"linkify/internal/services"
	"linkify/internal/workers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(c context.Context) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowCredentials: true,
	}))
	r.Use(middlewares.RateLimiter())

	api := r.Group("/api")

	cfg := config.Load()
	db, err := database.Connect(cfg.DSN)
	if err != nil {
		panic("Error: " + err.Error())
	}

	sr := repositories.NewSessionRepository(db)
	ss := services.NewSessionService(sr)

	sessionWorker := workers.NewSessionCleanupWorker(ss, time.Minute)

	go sessionWorker.Run(c)

	v1.RegisterRoutes(api.Group("/v1"), db)

	return r
}
