package models

import "github.com/golang-jwt/jwt/v4"

type Auth struct {
	Username  string `json:"username"`
	Aid       string `json:"aid"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
