package repositories

import (
	"context"
	"time"

	"linkify/internal/models"

	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Create(c context.Context, session *models.Session) error {
	return r.db.WithContext(c).Create(session).Error
}

func (r *SessionRepository) GetByTokenHash(c context.Context, tokenHash string) (*models.Session, error) {
	var session models.Session
	if err := r.db.WithContext(c).
		First(&session, "token_hash = ?", tokenHash).
		Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepository) DeleteByExpiresAt(c context.Context, now time.Time) error {
	return r.db.WithContext(c).Where("expires_at < ?", now).Delete(&models.Session{}).Error
}
