package repositories

import (
	"context"
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

func (r *BlockRepository) GetAll(c context.Context, profileID uuid.UUID) ([]models.Block, error) {
	var blocks []models.Block
	err := r.db.WithContext(c).Where("profile_id = ?", profileID).Find(&blocks).Error
	return blocks, err
}

func (r *BlockRepository) GetByID(c context.Context, id, userID, profileID uuid.UUID) (*models.Block, error) {
	var block *models.Block
	if err := r.db.WithContext(c).Where("id = ? AND user_id = ? AND profile_id = ?", id, userID, profileID).First(&block).Error; err != nil {
		return nil, err
	}

	return block, nil
}

func (r *BlockRepository) Update(c context.Context, block *models.Block) error {
	return r.db.WithContext(c).Save(block).Error
}

func (r *BlockRepository) Delete(c context.Context, id, userID, profileID uuid.UUID) error {
	result := r.db.WithContext(c).Delete(&models.Block{}, "id = ? AND user_id = ? AND profile_id = ?", id, userID, profileID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
