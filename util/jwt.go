package util

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
