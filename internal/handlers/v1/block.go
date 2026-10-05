package v1

import (
	"linkify/internal/dto"
	"linkify/internal/middlewares"
	"linkify/internal/repositories"
	"linkify/internal/services"
	"linkify/internal/uuidutil"
	"net/http"
	"uuid"

	"github.com/gin-gonic/gin"
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
