package repositories

import (
	"context"
	"uuid"

	"linkify/internal/models"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Create(c context.Context, profile *models.Profile) error {
	return r.db.WithContext(c).Create(profile).Error
}

func (r *ProfileRepository) GetAll(c context.Context, userID uuid.UUID) ([]models.Profile, error) {
	var profiles []models.Profile
	err := r.db.WithContext(c).Where("user_id = ?", userID).Find(&profiles).Error
	return profiles, err
}

func (r *ProfileRepository) GetByID(c context.Context, id uuid.UUID) (*models.Profile, error) {
	var profile models.Profile
	if err := r.db.WithContext(c).First(&profile, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *ProfileRepository) Update(c context.Context, profile *models.Profile) error {
	result := r.db.WithContext(c).Model(&models.Profile{}).Where("id = ?", profile.ID).Updates(profile)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ProfileRepository) Delete(c context.Context, id, userID uuid.UUID) error {
	result := r.db.WithContext(c).Where(&models.Profile{ID: id, UserID: userID}).Delete(&models.Profile{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
