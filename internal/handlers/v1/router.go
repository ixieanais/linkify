// Package v1
package v1

import (
	"linkify/internal/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	ur := repositories.NewUserRepository(db)
	RegisterAuthRoutes(r, ur)
	RegisterUserRoutes(r, ur)

	pr := repositories.NewProfileRepository(db)
	RegisterProfileRoutes(r, pr)

	lr := repositories.NewLinkRepository(db)
	RegisterLinkRoutes(r, lr)

	br := repositories.NewBlockRepository(db)
	RegisterBlockRoutes(r, br)
}
