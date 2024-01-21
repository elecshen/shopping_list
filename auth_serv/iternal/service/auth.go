package service

import (
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/elecshen/auth_service/iternal/model"
	"github.com/elecshen/auth_service/iternal/repository"
	"io"
	"strconv"
	"time"
)

const (
	PwSaltBytes = 64
	SingingKey  = "rgjef#4#%8GHNr43bj#rgek4FRMN"
	TokenTTL    = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) generateSalt(salt *[]byte) error {
	saltBytes := make([]byte, PwSaltBytes)
	_, err := io.ReadFull(rand.Reader, saltBytes)
	if err != nil {
		return err
	}
	*salt = saltBytes
	return nil
}

func (s *AuthService) generatePasswordHash(password string, salt []byte) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(salt))
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	if err := s.generateSalt(&user.Salt); err != nil {
		return 0, err
	}
	user.PasswordHash = s.generatePasswordHash(user.PasswordHash, user.Salt)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", err
	}

	if s.generatePasswordHash(password, user.Salt) != user.PasswordHash {
		return "", errors.New("access denied")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			Subject:   strconv.Itoa(user.Id),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(SingingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(SingingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return strconv.Atoi(claims.Subject)
}
