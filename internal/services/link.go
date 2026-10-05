package services

import (
	"context"
	"time"
	"uuid"

	"linkify/internal/models"
	"linkify/internal/repositories"
)

type LinkService struct {
	repo *repositories.LinkRepository
}

type CreateLinkInput struct {
	UserID    uuid.UUID
	ProfileID uuid.UUID
	URL       string
	Image     string
	IsActive  *bool
}

type UpdateLinkInput struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProfileID uuid.UUID
	URL       string
	Image     string
	IsActive  *bool
}

func NewLinkService(repo *repositories.LinkRepository) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) Create(c context.Context, input CreateLinkInput) (*models.Link, error) {
	now := time.Now()

	link := &models.Link{
		ID:        uuid.NewV4(),
		UserID:    input.UserID,
		ProfileID: input.ProfileID,
		URL:       input.URL,
		Image:     input.Image,
		IsActive:  input.IsActive,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(c, link); err != nil {
		return nil, err
	}

	return link, nil
}

func (s *LinkService) GetAll(c context.Context, profileID uuid.UUID) ([]models.Link, error) {
	return s.repo.GetAll(c, profileID)
}

func (s *LinkService) Get(c context.Context, id, userID, profileID uuid.UUID) (*models.Link, error) {
	return s.repo.GetByID(c, id, userID, profileID)
}

func (s *LinkService) Update(c context.Context, input UpdateLinkInput) error {
	link := &models.Link{
		ID:        input.ID,
		URL:       input.URL,
		Image:     input.Image,
		IsActive:  input.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.repo.Update(c, link)
}

func (s *LinkService) Delete(c context.Context, id, userID, profileID uuid.UUID) error {
	return s.repo.Delete(c, id, userID, profileID)
}
