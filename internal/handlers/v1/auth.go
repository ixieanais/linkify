package v1

import (
	"errors"
	"net/http"
	"time"

	"linkify/internal/dto"
	"linkify/internal/repositories"
	"linkify/internal/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Signup godoc
//
//	@Summary	Signup
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		payload	body	dto.UserRequest	true	"User Registration Details"
//	@Router		/auth/signup/ [post]
func (h *AuthHandler) Signup(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
		return
	}

	input := services.AuthInput{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.service.Signup(c, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	c.SetCookie(
		"session",
		*token,
		int(time.Now().Add(7*24*time.Hour).Unix()),
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusCreated, gin.H{
		"detail": "created",
	})
}

// Login godoc
//
//	@Summary	Login
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		payload	body	dto.UserRequest	true	"User Authentication Details"
//	@Router		/auth/login/ [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
	}

	input := services.AuthInput{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.service.Login(c, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) ||
			errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"detail": "invalid email or password",
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	c.SetCookie(
		"session",
		*token,
		int(time.Now().Add(7*24*time.Hour).Unix()),
		"/",
		"",
		false,
		true,
	)
}

func RegisterAuthRoutes(r *gin.RouterGroup, repo *repositories.UserRepository) {
	service := services.NewAuthService(repo)
	handler := NewAuthHandler(service)

	r.POST("/auth/signup/", handler.Signup)
	r.POST("/auth/login/", handler.Login)
}
