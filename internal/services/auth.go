package services

import (
	"context"
	"strings"
	"time"
	"uuid"

	"linkify/internal/models"
	"linkify/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	ur *repositories.UserRepository
	ss *SessionService
}

type AuthInput struct {
	Email    string
	Password string
}

func NewAuthService(
	ur *repositories.UserRepository,
	ss *SessionService,
) *AuthService {
	return &AuthService{ur: ur, ss: ss}
}

func (s *AuthService) Signup(c context.Context, input AuthInput) (*string, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	user := models.User{
		ID:        uuid.NewV4(),
		Email:     email,
		Password:  string(hash),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.ur.Create(c, &user); err != nil {
		return nil, err
	}

	session, err := s.ss.Create(c, user.ID)
	if err != nil {
		return nil, err
	}

	token, err := GenerateToken(user.ID, session.Token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *AuthService) Login(c context.Context, input AuthInput) (*string, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)

	user, err := s.ur.GetByEmail(c, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	session, err := s.ss.Create(c, user.ID)
	if err != nil {
		return nil, err
	}

	token, err := GenerateToken(user.ID, session.Token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
