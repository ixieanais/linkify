package repositories

import (
	"context"
	"errors"
	"uuid"

	"linkify/internal/models"

	"gorm.io/gorm"
)

type BlockRepository struct {
	db *gorm.DB
}

func NewBlockRepository(db *gorm.DB) *BlockRepository {
	return &BlockRepository{db: db}
}

func (r *BlockRepository) Create(c context.Context, block *models.Block) error {
	return r.db.WithContext(c).Create(block).Error
}

func (r *BlockRepository) GetAll(c context.Context) ([]models.Block, error) {
	var blocks []models.Block
	err := r.db.WithContext(c).Preload("Profile").Preload("Profile.User").Find(&blocks).Error
	return blocks, err
}

func (r *BlockRepository) GetByID(c context.Context, id uuid.UUID) (*models.Block, error) {
	var block *models.Block
	if err := r.db.WithContext(c).Preload("Profile").Preload("Profile.User").First(&block, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return block, nil
}

func (r *BlockRepository) Update(c context.Context, block *models.Block) error {
	return r.db.WithContext(c).Save(block).Error
}

func (r *BlockRepository) Delete(c context.Context, id uuid.UUID) error {
	return r.db.WithContext(c).Delete(&models.Block{}, id).Error
}
