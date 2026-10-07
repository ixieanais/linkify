// Package router
package router

import (
	"linkify/internal/config"
	"linkify/internal/database"
	v1 "linkify/internal/handlers/v1"
	"linkify/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middlewares.RateLimiter())

	api := r.Group("/api")

	cfg := config.Load()
	db, err := database.Connect(cfg.DSN)
	if err != nil {
		panic("Error: " + err.Error())
	}
	v1.RegisterRoutes(api.Group("/v1"), db)

	return r
}
