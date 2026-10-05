package v1

import (
	"errors"
	"net/http"

	"linkify/internal/dto"
	"linkify/internal/middlewares"
	"linkify/internal/repositories"
	"linkify/internal/services"
	"linkify/internal/uuidutil"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Get godoc
//
//	@Summary	Get user
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{object}	dto.UserResponse
//	@Router		/users/me/ [get]
func (h *UserHandler) Get(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "user id not found in context",
		})
		return
	}

	userID, err := uuidutil.ParseUUID(rawUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
		return
	}

	user, err := h.service.Get(c, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"detail": "user not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
		return
	}

	response := dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// Update godoc
//
//	@Summary	Update User
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		payload	body	dto.UserRequest	true	"User Updation Details"
//	@Router		/users/me/ [put]
func (h *UserHandler) Update(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
		return
	}

	rawUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "user id not found in context",
		})
		return
	}

	userID, err := uuidutil.ParseUUID(rawUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	input := services.UpdateUserInput{
		ID:       userID,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.service.Update(c, input); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"detail": "user not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}
}

// Delete godoc
//
//	@Summary	Delete User
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Router		/users/me/ [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "user id not found in context",
		})
		return
	}

	userID, err := uuidutil.ParseUUID(rawUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	if err := h.service.Delete(c, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"detail": "user not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}
}

func RegisterUserRoutes(r *gin.RouterGroup, repo *repositories.UserRepository) {
	service := services.NewUserService(repo)
	handler := NewUserHandler(service)

	r.Use(middlewares.AuthMiddleware())
	r.GET("/users/me/", handler.Get)
	r.PUT("/users/me/", handler.Update)
	r.DELETE("/users/me/", handler.Delete)
}
