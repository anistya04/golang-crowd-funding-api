package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type Service interface {
	GenerateToken(userID uint64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct{}

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID uint64) (string, error) {

	claims := jwt.MapClaims{}
	claims["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("jws")))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("jws")), nil
	})

	if err != nil {
		return parsedToken, err
	}

	return parsedToken, nil

}
