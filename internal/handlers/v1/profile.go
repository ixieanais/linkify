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
//	@Param		payload	body		dto.ProfileRequest	true	"User Creation Details"
//	@Success	200		{object}	dto.ProfileResponse
//	@Security	ApiKeyAuth
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
		return
	}

	input := services.CreateProfileInput{
		UserID:    userID,
		Username:  request.Username,
		Name:      request.Name,
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
		Name:      profile.Name,
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
//	@Success	200	{array}	dto.ProfileResponse
//	@Security	ApiKeyAuth
//	@Router		/profiles/ [get]
func (h *ProfileHandler) GetAll(c *gin.Context) {
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

	profiles, err := h.service.GetAll(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
		return
	}

	response := []dto.ProfileResponse{}

	for _, profile := range profiles {
		response = append(response, dto.ProfileResponse{
			ID:        profile.ID,
			Username:  profile.Username,
			Name:      profile.Name,
			Bio:       profile.Bio,
			AvatarURL: profile.AvatarURL,
			CreatedAt: profile.CreatedAt,
			UpdatedAt: profile.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
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
		Name:      profile.Name,
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
//	@Param		id		path	string				true	"Profile ID"	format(uuid)
//	@Param		payload	body	dto.ProfileRequest	true	"Profile Updation Details"
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/ [put]
func (h *ProfileHandler) Update(c *gin.Context) {
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

	input := services.UpdateProfileInput{
		ID:        id,
		UserID:    userID,
		Username:  req.Username,
		Name:      req.Name,
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
		return
	}
}

// Delete godoc
//
//	@Summary	Delete profile
//	@Tags		profiles
//	@Accept		json
//	@Produce	json
//	@Param		id	path	string	true	"Profile ID"	format(uuid)
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/ [delete]
func (h *ProfileHandler) Delete(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": "user id not found in context",
		})
	}

	userID, err := uuidutil.ParseUUID(rawUserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
	}

	rawID := c.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
	}

	if err := h.service.Delete(c, id, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"detail": "profile not found",
			})
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
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
