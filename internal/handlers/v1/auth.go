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

	c.SetCookieData(&http.Cookie{
		Name:     "session_id",
		Value:    *token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(7 * 24 * time.Hour),
	})

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
		return
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
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	c.SetCookieData(&http.Cookie{
		Name:     "session_id",
		Value:    *token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(7 * 24 * time.Hour),
	})

	c.JSON(http.StatusOK, gin.H{
		"detail": "logged in",
	})
}

func RegisterAuthRoutes(
	r *gin.RouterGroup,
	ur *repositories.UserRepository,
	ss *services.SessionService,
) {
	service := services.NewAuthService(ur, ss)
	handler := NewAuthHandler(service)

	r.POST("/auth/signup/", handler.Signup)
	r.POST("/auth/login/", handler.Login)
}
