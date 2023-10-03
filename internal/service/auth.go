package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"net/http"
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/config"
	"otus-social-network/internal/dto"
	"otus-social-network/internal/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const invalidTokenMessage = "invalid token"

type AuthServiceInstance struct {
	tokenTTL          uint
	signingKey        string
	salt              string
	repositoryManager *repository.RepositoryManager
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"user_id`
}

type AuthService interface {
	CreateUser(user *dto.SignUpInput) (uuid.UUID, *app_error.HttpError)
	GenerateToken(email, password string) (string, *app_error.HttpError)
	ParseToken(accessToken string) (uuid.UUID, *app_error.HttpError)
}

func NewAuthService(repositoryManager *repository.RepositoryManager, config *config.Config) *AuthServiceInstance {
	return &AuthServiceInstance{
		salt:              config.Salt,
		signingKey:        config.SigningKey,
		repositoryManager: repositoryManager,
		tokenTTL:          config.TokenTTL,
	}
}

func (s *AuthServiceInstance) GenerateToken(email, password string) (string, *app_error.HttpError) {
	user, err := s.repositoryManager.UserRepository.GetUser(email, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	expireAt := time.Now().Add(time.Second * time.Duration(s.tokenTTL))
	IssuedAt := time.Now().Unix()

	logrus.Debugln(expireAt)
	logrus.Debugln(IssuedAt)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(s.tokenTTL)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	tokenString, erro := token.SignedString([]byte(s.signingKey))

	if erro != nil {
		return "", app_error.NewInternalServerError(erro)
	}

	return tokenString, nil
}

func (s *AuthServiceInstance) ParseToken(accessToken string) (uuid.UUID, *app_error.HttpError) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return uuid.Nil, app_error.NewHttpError(errors.New("invalid signing method"), invalidTokenMessage, "id", http.StatusBadRequest)
		}

		return []byte(s.signingKey), nil
	})

	if err != nil {
		return uuid.Nil, app_error.NewHttpError(err, invalidTokenMessage, "id", http.StatusBadRequest)
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return uuid.Nil, app_error.NewHttpError(errors.New("token claims are not of type tokenClaims"), invalidTokenMessage, "id", http.StatusBadRequest)
	}

	return claims.UserId, nil
}

func (s *AuthServiceInstance) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt)))
}

func (s *AuthServiceInstance) CreateUser(user *dto.SignUpInput) (uuid.UUID, *app_error.HttpError) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repositoryManager.CreateUser(user)
}
