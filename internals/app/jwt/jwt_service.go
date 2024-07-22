package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aparnasukesh/auth-svc/pkg/common"
	"github.com/golang-jwt/jwt"
)

type service struct {
	jwt_secret_key string
}

func NewJWTService(jwt_secret_key string) common.JWT_Service {
	return &service{
		jwt_secret_key: jwt_secret_key,
	}
}

func (s *service) GenerateJWT(email string, userId, roleId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["role"] = roleId
	claims["email"] = email
	claims["userid"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()

	secretKey := []byte(s.jwt_secret_key)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *service) VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_secret_key")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *service) GetRole(token *jwt.Token) (interface{}, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	roleId := claims["role"]
	return roleId, nil
}

func (s *service) GetUserID(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims type")
	}

	userID, ok := claims["userid"].(float64)
	if !ok {
		return 0, errors.New("user ID not found or not a number in claims")
	}

	return int(userID), nil
}
