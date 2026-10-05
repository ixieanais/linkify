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

type BlockHandler struct {
	service *services.BlockService
}

func NewBlockHandler(service *services.BlockService) *BlockHandler {
	return &BlockHandler{service: service}
}

// Create godoc
//
//	@Summary	Create block
//	@Tags		blocks
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string				true	"Profile ID"	format(uuid)
//	@Param		payload	body		dto.BlockRequest	true	"Block Creation Details"
//	@Success	201		{object}	dto.BlockResponse
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/blocks/ [post]
func (h *BlockHandler) Create(c *gin.Context) {
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

	var req dto.BlockRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
	}

	input := services.CreateBlockInput{
		UserID:    userID,
		ProfileID: profileID,
		Title:     req.Title,
		Text:      req.Text,
		Image:     req.Image,
		URL:       req.URL,
		Type:      req.Type,
		IsActive:  req.IsActive,
	}

	block, err := h.service.Create(c, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	response := dto.BlockResponse{
		ID:        block.ID,
		Title:     block.Title,
		Text:      block.Text,
		Image:     block.Image,
		URL:       block.URL,
		Type:      block.Type,
		IsActive:  block.IsActive,
		CreatedAt: block.CreatedAt,
		UpdatedAt: block.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAll godoc
//
//	@Summary	Get blocks
//	@Tags		blocks
//	@Accept		json
//	@Produce	json
//	@Param		id	path	string	true	"Profile ID"	format(uuid)
//	@Success	200	{array}	dto.BlockResponse
//	@Router		/profiles/{id}/blocks/ [get]
func (h *BlockHandler) GetAll(c *gin.Context) {
	rawProfileID := c.Param("id")
	profileID, err := uuid.Parse(rawProfileID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid profile id type",
		})
	}

	blocks, err := h.service.GetAll(c, profileID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	response := []dto.BlockResponse{}

	for _, block := range blocks {
		response = append(response, dto.BlockResponse{
			ID:        block.ID,
			Title:     block.Title,
			Text:      block.Text,
			Image:     block.Image,
			URL:       block.URL,
			Type:      block.Type,
			IsActive:  block.IsActive,
			CreatedAt: block.CreatedAt,
			UpdatedAt: block.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

// Get godoc
//
//	@Summary	Get block
//	@Tags		blocks
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string	true	"Profile ID"	format(uuid)
//	@Param		blockId	path		string	true	"Block ID"		format(uuid)
//	@Success	200		{object}	dto.BlockResponse
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/blocks/{blockId}/ [get]
func (h *BlockHandler) Get(c *gin.Context) {
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

	rawBlockID := c.Param("bid")
	blockID, err := uuid.Parse(rawBlockID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid block id type",
		})
	}

	block, err := h.service.Get(c, blockID, userID, profileID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"detail": "block not found",
			})
		}
	}

	response := dto.BlockResponse{
		ID:        block.ID,
		Title:     block.Title,
		Text:      block.Text,
		Image:     block.Image,
		URL:       block.URL,
		Type:      block.Type,
		IsActive:  block.IsActive,
		CreatedAt: block.CreatedAt,
		UpdatedAt: block.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// Update godoc
//
//	@Summary	Update block
//	@Tags		blocks
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string				true	"Profile ID"	format(uuid)
//	@Param		blockId	path	string				true	"Block ID"		format(uuid)
//	@Param		payload	body	dto.BlockRequest	true	"Block Updation Details"
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/blocks/{blockId}/ [put]
func (h *BlockHandler) Update(c *gin.Context) {
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

	rawBlockID := c.Param("bid")
	blockID, err := uuid.Parse(rawBlockID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid block id type",
		})
	}

	var req dto.BlockRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
	}

	input := services.UpdateBlockInput{
		ID:        blockID,
		UserID:    userID,
		ProfileID: profileID,
		Title:     req.Title,
		Text:      req.Text,
		Image:     req.Image,
		URL:       req.URL,
		Type:      req.Type,
		IsActive:  req.IsActive,
	}

	if err := h.service.Update(c, input); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"detail": "block not found",
			})
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"detail": "updated",
	})
}

// Delete godoc
//
//	@Summary	Delete block
//	@Tags		blocks
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string	true	"Profile ID"	format(uuid)
//	@Param		blockId	path	string	true	"Block ID"		format(uuid)
//	@Security	ApiKeyAuth
//	@Router		/profiles/{id}/blocks/{blockId}/ [delete]
func (h *BlockHandler) Delete(c *gin.Context) {
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

	rawBlockID := c.Param("bid")
	blockID, err := uuid.Parse(rawBlockID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "invalid block id type",
		})
	}

	if err := h.service.Delete(c, blockID, userID, profileID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"detail": "block not found",
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

func RegisterBlockRoutes(r *gin.RouterGroup, repo *repositories.BlockRepository) {
	service := services.NewBlockService(repo)
	handler := NewBlockHandler(service)

	r.POST("/profiles/:id/blocks/", middlewares.AuthMiddleware(), handler.Create)
	r.GET("/profiles/:id/blocks/", handler.GetAll)
	r.GET("/profiles/:id/blocks/:bid/", middlewares.AuthMiddleware(), handler.Get)
	r.PUT("/profiles/:id/blocks/:bid/", middlewares.AuthMiddleware(), handler.Update)
	r.DELETE("/profiles/:id/blocks/:bid/", middlewares.AuthMiddleware(), handler.Delete)
}
