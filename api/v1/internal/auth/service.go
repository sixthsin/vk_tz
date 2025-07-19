package auth

import (
	"errors"
	"marketplace-api/config"
	"marketplace-api/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceDeps struct {
	*config.Config
	*user.UserRepository
}

type AuthService struct {
	*config.Config
	*user.UserRepository
}

func NewAuthService(deps AuthServiceDeps) *AuthService {
	return &AuthService{
		Config:         deps.Config,
		UserRepository: deps.UserRepository,
	}
}

func (s *AuthService) Register(username, password string) (string, error) {
	existedUsername, err := s.UserRepository.FindByUsername(username)
	if err != nil {
		return "", nil
	}
	if existedUsername != nil {
		return "", errors.New(ErrUsernameExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Username: username,
		Password: string(hashedPassword),
	}

	createdUser, err := s.UserRepository.Create(user)
	if err != nil {
		return "", err
	}

	return createdUser.Username, nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	existedUser, err := s.UserRepository.FindByUsername(username)
	if err != nil {
		return "", nil
	}
	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}

	return existedUser.Username, nil
}
