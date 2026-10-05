package repositories

import (
	"context"
	"uuid"

	"linkify/internal/models"

	"gorm.io/gorm"
)

type LinkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(c context.Context, link *models.Link) error {
	return r.db.WithContext(c).Create(link).Error
}

func (r *LinkRepository) GetAll(c context.Context, profileID uuid.UUID) ([]models.Link, error) {
	var links []models.Link
	err := r.db.WithContext(c).Where("profile_id = ?", profileID).Find(&links).Error
	return links, err
}

func (r *LinkRepository) GetByID(c context.Context, id, userID, profileID uuid.UUID) (*models.Link, error) {
	var link *models.Link
	if err := r.db.WithContext(c).Where("id = ? AND user_id = ? AND profile_id = ?", id, userID, profileID).First(link).Error; err != nil {
		return nil, err
	}

	return link, nil
}

func (r *LinkRepository) Update(c context.Context, link *models.Link) error {
	result := r.db.WithContext(c).Model(&models.Link{}).Where("id = ?", link.ID).Updates(link)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *LinkRepository) Delete(c context.Context, id, userID, profileID uuid.UUID) error {
	result := r.db.WithContext(c).Delete(&models.Link{}, "id = ? AND user_id = ? AND profile_id = ?", id, userID, profileID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
