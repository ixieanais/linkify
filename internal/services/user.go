// Package services
package services

import (
	"context"
	"errors"
	"strings"
	"time"
	"uuid"

	"linkify/internal/models"
	"linkify/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

type CreateUserInput struct {
	Email    string
	Password string
}

type UpdateUserInput struct {
	ID       uuid.UUID
	Email    string
	Password string
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(c context.Context, input CreateUserInput) (*models.User, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)

	if email == "" {
		return nil, errors.New("email is required")
	}

	if password == "" {
		return nil, errors.New("password is required")
	}

	now := time.Now()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:        uuid.NewV4(),
		Email:     email,
		Password:  string(hash),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(c, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAll(c context.Context) ([]models.User, error) {
	return s.repo.GetAll(c)
}

func (s *UserService) Get(c context.Context, id uuid.UUID) (*models.User, error) {
	return s.repo.GetByID(c, id)
}

func (s *UserService) Update(c context.Context, input UpdateUserInput) error {
	email := strings.ToLower(strings.TrimSpace(input.Email))

	if email == "" {
		return errors.New("email is required")
	}

	user := &models.User{
		ID:        input.ID,
		Email:     email,
		UpdatedAt: time.Now(),
	}

	return s.repo.Update(c, user)
}

func (s *UserService) Delete(c context.Context, id uuid.UUID) error {
	return s.repo.Delete(c, id)
}
