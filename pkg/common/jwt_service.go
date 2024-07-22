package common

import "github.com/golang-jwt/jwt"

type JWT_Service interface {
	GenerateJWT(email string, userId, roleId uint) (string, error)
	VerifyJWT(token string) (*jwt.Token, error)
	GetRole(token *jwt.Token) (interface{}, error)
	GetUserID(token *jwt.Token) (int, error)
}
