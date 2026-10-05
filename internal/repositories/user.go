// Package repositories
package repositories

import (
	"context"
	"uuid"

	"linkify/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(c context.Context, user *models.User) error {
	return r.db.WithContext(c).Create(user).Error
}

func (r *UserRepository) GetAll(c context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(c).Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByID(c context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(c).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(c context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(c).First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(c context.Context, user *models.User) error {
	result := r.db.WithContext(c).Model(&models.User{}).Where("id = ?", user.ID).Updates(user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *UserRepository) Delete(c context.Context, id uuid.UUID) error {
	result := r.db.WithContext(c).Delete(&models.User{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
