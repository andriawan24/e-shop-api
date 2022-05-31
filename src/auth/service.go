package auth

import (
	"e-shop/src/utils"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authService struct{}

func NewService() *authService {
	return &authService{}
}

func (s *authService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["users_id"] = userID
	ttl := 60 * time.Hour
	claim["exp"] = time.Now().UTC().Add(ttl).Unix()

	secretKey, err := utils.GetSecretKey()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *authService) ValidateToken(tokenEncoded string) (*jwt.Token, error) {
	secretKey, err := utils.GetSecretKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenEncoded, func(tokenEncoded *jwt.Token) (interface{}, error) {
		_, ok := tokenEncoded.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid credentials")
		}

		return secretKey, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil

}
