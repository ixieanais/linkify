package v1

import (
	"linkify/internal/repositories"
	"linkify/internal/services"

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
//	@Router		/blocks/ [post]
func (h *BlockHandler) Create(c *gin.Context) {
}

// GetAll godoc
//
//	@Summary	Get blocks
//	@Tags		blocks
//	@Accept		json
//	@Produce	json
//	@Router		/blocks/ [get]
func (h *BlockHandler) GetAll(c *gin.Context) {
}

// Get godoc
//
//	@Summary	Get block
//	@Tags		blocks
//	@Accept		json
//	@Produce	json
//	@Router		/blocks/{id}/ [get]
func (h *BlockHandler) Get(c *gin.Context) {
}

// Update godoc
//
//	@Summary	Update block
//	@Tags		blocks
//	@Accept		json
//	@Produce	json
//	@Router		/blocks/{id}/ [put]
func (h *BlockHandler) Update(c *gin.Context) {
}

// Delete godoc
//
//	@Summary	Delete block
//	@Tags		blocks
//	@Accept		json
//	@Produce	json
//	@Router		/blocks/{id}/ [delete]
func (h *BlockHandler) Delete(c *gin.Context) {
}

func RegisterBlockRoutes(r *gin.RouterGroup, repo *repositories.BlockRepository) {
	service := services.NewBlockService(repo)
	handler := NewBlockHandler(service)

	r.POST("/blocks/", handler.Create)
	r.GET("/blocks/", handler.GetAll)
	r.GET("/blocks/:id/", handler.Get)
	r.PUT("/blocks/:id/", handler.Update)
	r.DELETE("/blocks/:id/", handler.Delete)
}
