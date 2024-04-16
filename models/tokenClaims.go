package models

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
	Iat int64  `json:"iat"`
	jwt.RegisteredClaims
}
