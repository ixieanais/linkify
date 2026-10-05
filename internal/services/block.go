package services

import (
	"context"
	"time"
	"uuid"

	"linkify/internal/models"
	"linkify/internal/repositories"
)

type BlockService struct {
	repo *repositories.BlockRepository
}

type CreateBlockInput struct {
	UserID    uuid.UUID
	ProfileID uuid.UUID
	Title     string
	Text      *string
	Image     *string
	URL       *string
	Type      string
	IsActive  *bool
}

type UpdateBlockInput struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProfileID uuid.UUID
	Title     string
	Text      *string
	Image     *string
	URL       *string
	Type      string
	IsActive  *bool
}

func NewBlockService(repo *repositories.BlockRepository) *BlockService {
	return &BlockService{repo: repo}
}

func (s *BlockService) Create(c context.Context, input CreateBlockInput) (*models.Block, error) {
	now := time.Now()

	block := &models.Block{
		ID:        uuid.NewV4(),
		UserID:    input.UserID,
		ProfileID: input.ProfileID,
		Title:     input.Title,
		Text:      input.Text,
		Image:     input.Image,
		URL:       input.URL,
		Type:      input.Type,
		IsActive:  input.IsActive,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(c, block); err != nil {
		return nil, err
	}

	return block, nil
}

func (s *BlockService) GetAll(c context.Context, profileID uuid.UUID) ([]models.Block, error) {
	return s.repo.GetAll(c, profileID)
}

func (s *BlockService) Get(c context.Context, id, userID, profileID uuid.UUID) (*models.Block, error) {
	return s.repo.GetByID(c, id, userID, profileID)
}

func (s *BlockService) Update(c context.Context, input UpdateBlockInput) error {
	block := &models.Block{
		ID:        input.ID,
		Title:     input.Title,
		Text:      input.Text,
		Image:     input.Image,
		URL:       input.URL,
		Type:      input.Type,
		IsActive:  input.IsActive,
		UpdatedAt: time.Now(),
	}

	return s.repo.Update(c, block)
}

func (s *BlockService) Delete(c context.Context, id, userID, profileID uuid.UUID) error {
	return s.repo.Delete(c, id, userID, profileID)
}
