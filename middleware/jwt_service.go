package middleware

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey []byte
}

func NewJWTService() *jwtService {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("SECRET KEY is not set in environment variables")
	}

	return &jwtService{
		secretKey: []byte(secret),
	}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	var secretKey = []byte(s.secretKey)

	claim := jwt.MapClaims{
		"id": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(s.secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}