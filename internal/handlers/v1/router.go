// Package v1
package v1

import (
	"linkify/internal/repositories"
	"linkify/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	ur := repositories.NewUserRepository(db)
	sr := repositories.NewSessionRepository(db)
	ss := services.NewSessionService(sr)
	RegisterAuthRoutes(r, ur, ss)
	RegisterUserRoutes(r, ur, ss)

	pr := repositories.NewProfileRepository(db)
	RegisterProfileRoutes(r, pr, ss)

	lr := repositories.NewLinkRepository(db)
	RegisterLinkRoutes(r, lr, ss)

	br := repositories.NewBlockRepository(db)
	RegisterBlockRoutes(r, br, ss)
}
