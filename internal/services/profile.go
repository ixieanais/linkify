package services

import (
	"context"
	"strings"
	"time"
	"uuid"

	"linkify/internal/models"
	"linkify/internal/repositories"
)

type ProfileService struct {
	repo *repositories.ProfileRepository
}

type CreateProfileInput struct {
	UserID     uuid.UUID
	Slug       string
	Name       string
	Bio        *string
	AvatarURL  *string
	Background string
}

type UpdateProfileInput struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Slug       string
	Name       string
	Bio        *string
	AvatarURL  *string
	Background string
}

func NewProfileService(repo *repositories.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) Create(c context.Context, input CreateProfileInput) (*models.Profile, error) {
	slug := strings.TrimSpace(input.Slug)
	name := strings.TrimSpace(input.Name)

	now := time.Now()

	profile := &models.Profile{
		ID:         uuid.NewV4(),
		UserID:     input.UserID,
		Slug:       slug,
		Name:       name,
		Bio:        input.Bio,
		AvatarURL:  input.AvatarURL,
		Background: input.Background,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if err := s.repo.Create(c, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *ProfileService) GetAll(c context.Context, userID uuid.UUID) ([]models.Profile, error) {
	return s.repo.GetAll(c, userID)
}

func (s *ProfileService) Get(c context.Context, id uuid.UUID) (*models.Profile, error) {
	return s.repo.GetByID(c, id)
}

func (s *ProfileService) GetByUsername(c context.Context, slug string) (*models.Profile, error) {
	return s.repo.GetByUsername(c, slug)
}

func (s *ProfileService) Update(c context.Context, input UpdateProfileInput) error {
	slug := strings.TrimSpace(input.Slug)
	name := strings.TrimSpace(input.Name)

	profile := &models.Profile{
		ID:         input.ID,
		UserID:     input.UserID,
		Slug:       slug,
		Name:       name,
		Bio:        input.Bio,
		AvatarURL:  input.AvatarURL,
		Background: input.Background,
		UpdatedAt:  time.Now(),
	}

	return s.repo.Update(c, profile)
}

func (s *ProfileService) Delete(c context.Context, id, userID uuid.UUID) error {
	return s.repo.Delete(c, id, userID)
}
