package services

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"time"
	"uuid"

	"linkify/internal/models"
	"linkify/internal/repositories"
)

type SessionService struct {
	repo *repositories.SessionRepository
}

type SessionResponse struct {
	Token     string
	ExpiresAt time.Time
}

func NewSessionService(repo *repositories.SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

func (s *SessionService) Create(c context.Context, userID uuid.UUID) (*SessionResponse, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return nil, err
	}

	token := base64.RawURLEncoding.EncodeToString(tokenBytes)
	rawTokenHash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(rawTokenHash[:])

	now := time.Now()
	expiresAt := now.Add(7 * 24 * time.Hour)

	session := &models.Session{
		ID:        uuid.NewV4(),
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: expiresAt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(c, session); err != nil {
		return nil, err
	}

	response := &SessionResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}

	return response, nil
}

func (s *SessionService) Validate(c context.Context, token string) (*models.Session, error) {
	rawTokenHash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(rawTokenHash[:])

	session, err := s.repo.GetByTokenHash(c, tokenHash)
	if err != nil {
		return nil, err
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, errors.New("session expired")
	}

	return session, nil
}

func (s *SessionService) DeleteByExpiresAt(c context.Context) error {
	now := time.Now()
	return s.repo.DeleteByExpiresAt(c, now)
}
