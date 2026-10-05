package v1

import (
	"errors"
	"net/http"
	"uuid"

	"linkify/internal/dto"
	"linkify/internal/middlewares"
	"linkify/internal/repositories"
	"linkify/internal/services"
	"linkify/internal/uuidutil"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProfileHandler struct {
	service *services.ProfileService
}

func NewProfileHandler(service *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

// Create godoc
//
//	@Summary	Create profile
//	@Tags		profiles
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		payload	body		dto.ProfileRequest	true	"User Creation Details"
//	@Success	200		{object}	dto.ProfileResponse
//	@Router		/profiles/ [post]
func (h *ProfileHandler) Create(c *gin.Context) {
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

	var request dto.ProfileRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": "invalid profile struct",
		})
	}

	input := services.CreateProfileInput{
		UserID:    userID,
		Username:  request.Username,
		URL:       request.URL,
		Bio:       request.Bio,
		AvatarURL: request.AvatarURL,
	}

	profile, err := h.service.Create(c, input)
	if err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "invalid user id",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
		return
	}

	response := dto.ProfileResponse{
		ID:        profile.ID,
		Username:  profile.Username,
		URL:       profile.URL,
		Bio:       profile.Bio,
		AvatarURL: profile.AvatarURL,
		CreatedAt: profile.CreatedAt,
		UpdatedAt: profile.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAll godoc
//
//	@Summary	Get profiles
//	@Tags		profiles
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Router		/profiles/ [get]
func (h *ProfileHandler) GetAll(c *gin.Context) {
}

// Get godoc
//
//	@Summary	Get profile
//	@Tags		profiles
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Profile ID"	format(uuid)
//	@Success	200	{object}	dto.ProfileResponse
//	@Router		/profiles/{id}/ [get]
func (h *ProfileHandler) Get(c *gin.Context) {
	rawID := c.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	profile, err := h.service.Get(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"detail": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
		return
	}

	response := dto.ProfileResponse{
		ID:        profile.ID,
		Username:  profile.Username,
		URL:       profile.URL,
		Bio:       profile.Bio,
		AvatarURL: profile.AvatarURL,
		CreatedAt: profile.CreatedAt,
		UpdatedAt: profile.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// Update godoc
//
//	@Summary	Update profile
//	@Tags		profiles
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		id		path	string				true	"Profile ID"	format(uuid)
//	@Param		payload	body	dto.ProfileRequest	true	"Profile Updation Details"
//	@Router		/profiles/{id}/ [put]
func (h *ProfileHandler) Update(c *gin.Context) {
	rawID := c.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	var req dto.ProfileRequest

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

	input := services.UpdateProfileInput{
		ID:        id,
		UserID:    userID,
		Username:  req.Username,
		URL:       req.URL,
		Bio:       req.Bio,
		AvatarURL: req.AvatarURL,
	}

	if err := h.service.Update(c, input); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"detail": "profile not found",
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
//	@Summary	Delete profile
//	@Tags		profiles
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		id	path	string	true	"Profile ID"	format(uuid)
//	@Router		/profiles/{id}/ [delete]
func (h *ProfileHandler) Delete(c *gin.Context) {
	rawID := c.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

	if err := h.service.Delete(c, id, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"detail": "profile not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}
}

func RegisterProfileRoutes(r *gin.RouterGroup, repo *repositories.ProfileRepository) {
	service := services.NewProfileService(repo)
	handler := NewProfileHandler(service)

	r.POST("/profiles/", middlewares.AuthMiddleware(), handler.Create)
	r.GET("/profiles/", middlewares.AuthMiddleware(), handler.GetAll)
	r.GET("/profiles/:id/", handler.Get)
	r.PUT("/profiles/:id/", middlewares.AuthMiddleware(), handler.Update)
	r.DELETE("/profiles/:id/", middlewares.AuthMiddleware(), handler.Delete)
}
