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

type LinkHandler struct {
	service *services.LinkService
}

func NewLinkHandler(service *services.LinkService) *LinkHandler {
	return &LinkHandler{service: service}
}

// Create godoc
//
//	@Summary	Create link
//	@Tags		links
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string			true	"Profile ID"	format(uuid)
//	@Param		payload	body		dto.LinkRequest	true	"Profile Creation Details"
//	@Success	201		{object}	dto.LinkResponse
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/links/ [post]
func (h *LinkHandler) Create(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": "user id not found in context",
		})
	}

	userID, err := uuidutil.ParseUUID(rawUserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	rawProfileID := c.Param("id")
	profileID, err := uuid.Parse(rawProfileID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid profile id type",
		})
	}

	var req dto.LinkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
	}

	input := services.CreateLinkInput{
		UserID:    userID,
		ProfileID: profileID,
		URL:       req.URL,
		Image:     req.Image,
		IsActive:  req.IsActive,
	}

	link, err := h.service.Create(c, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	response := dto.LinkResponse{
		ID:        link.ID,
		URL:       link.URL,
		Image:     link.Image,
		IsActive:  link.IsActive,
		CreatedAt: link.CreatedAt,
		UpdatedAt: link.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAll godoc
//
//	@Summary	Get links
//	@Tags		links
//	@Accept		json
//	@Produce	json
//	@Param		id	path	string	true	"Profile ID"	format(uuid)
//	@Success	200	{array}	dto.LinkResponse
//	@Router		/profiles/{id}/links/ [get]
func (h *LinkHandler) GetAll(c *gin.Context) {
	rawProfileID := c.Param("id")
	profileID, err := uuid.Parse(rawProfileID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid profile id type",
		})
	}

	links, err := h.service.GetAll(c, profileID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	response := []dto.LinkResponse{}

	for _, link := range links {
		response = append(response, dto.LinkResponse{
			ID:        link.ID,
			URL:       link.URL,
			Image:     link.Image,
			IsActive:  link.IsActive,
			CreatedAt: link.CreatedAt,
			UpdatedAt: link.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

// Get godoc
//
//	@Summary	Get link
//	@Tags		links
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string	true	"Profile ID"	format(uuid)
//	@Param		linkId	path		string	true	"Link ID"		format(uuid)
//	@Success	200		{object}	dto.LinkResponse
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/links/{linkId} [get]
func (h *LinkHandler) Get(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": "user id not found in context",
		})
	}
	userID, err := uuidutil.ParseUUID(rawUserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	rawProfileID := c.Param("id")
	profileID, err := uuid.Parse(rawProfileID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid profile id type",
		})
	}

	rawLinkID := c.Param("lid")
	linkID, err := uuid.Parse(rawLinkID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid link id type",
		})
	}

	link, err := h.service.Get(c, linkID, userID, profileID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"detail": "link not found",
			})
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	response := dto.LinkResponse{
		ID:        link.ID,
		URL:       link.URL,
		Image:     link.Image,
		IsActive:  link.IsActive,
		CreatedAt: link.CreatedAt,
		UpdatedAt: link.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// Update godoc
//
//	@Summary	Update link
//	@Tags		links
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string			true	"Profile ID"	format(uuid)
//	@Param		linkId	path	string			true	"Link ID"		format(uuid)
//	@Param		payload	body	dto.LinkRequest	true	"Link Updation Details"
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/links/{linkId} [put]
func (h *LinkHandler) Update(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": "user id not found in context",
		})
	}

	userID, err := uuidutil.ParseUUID(rawUserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	rawProfileID := c.Param("id")
	profileID, err := uuid.Parse(rawProfileID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid profile id type",
		})
	}

	rawLinkID := c.Param("lid")
	linkID, err := uuid.Parse(rawLinkID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid link id type",
		})
	}

	var req dto.LinkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
	}

	input := services.UpdateLinkInput{
		ID:        linkID,
		UserID:    userID,
		ProfileID: profileID,
		URL:       req.URL,
		Image:     req.Image,
		IsActive:  req.IsActive,
	}

	if err := h.service.Update(c, input); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"detail": "link not found",
			})
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"detail": "edited",
	})
}

// Delete godoc
//
//	@Summary	Delete link
//	@Tags		links
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string	true	"Profile ID"	format(uuid)
//	@Param		linkId	path	string	true	"Link ID"		format(uuid)
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/links/{linkId}/ [delete]
func (h *LinkHandler) Delete(c *gin.Context) {
	rawUserID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": "user id not found in context",
		})
	}

	userID, err := uuidutil.ParseUUID(rawUserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	rawProfileID := c.Param("id")
	profileID, err := uuid.Parse(rawProfileID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid profile id type",
		})
	}

	rawLinkID := c.Param("lid")
	linkID, err := uuid.Parse(rawLinkID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid link id type",
		})
	}

	if err := h.service.Delete(c, linkID, userID, profileID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"detail": "link not found",
			})
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"detail": "deleted",
	})
}

func RegisterLinkRoutes(r *gin.RouterGroup, repo *repositories.LinkRepository) {
	service := services.NewLinkService(repo)
	handler := NewLinkHandler(service)

	r.POST("/profiles/:id/links/", middlewares.AuthMiddleware(), handler.Create)
	r.GET("/profiles/:id/links/", handler.GetAll)
	r.GET("/profiles/:id/links/:lid/", middlewares.AuthMiddleware(), handler.Get)
	r.PUT("/profiles/:id/links/:lid/", middlewares.AuthMiddleware(), handler.Update)
	r.DELETE("/profiles/:id/links/:lid/", middlewares.AuthMiddleware(), handler.Delete)
}
