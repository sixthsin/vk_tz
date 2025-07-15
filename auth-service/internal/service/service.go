package service

import (
	"auth-service/config"
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceDeps struct {
	*config.Config
	*repository.UserRepository
}

type AuthService struct {
	*config.Config
	*repository.UserRepository
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

	user := &model.User{
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

func (s *AuthService) ValidateToken(tokenString string) (*UserResponse, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(s.Config.Auth.Secret), nil
	})
	if err != nil || !token.Valid {
		return &UserResponse{IsValid: false}, errors.New(ErrInvalidToken)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &UserResponse{IsValid: false}, errors.New(ErrInvalidTokenClaims)
	}

	expirationTime := int64(claims["exp"].(float64))
	if time.Now().Unix() > expirationTime {
		return &UserResponse{IsValid: false}, errors.New(ErrTokenExpired)
	}

	username := claims["username"].(string)

	return &UserResponse{
		Username: username,
		IsValid:  true,
	}, nil
}
