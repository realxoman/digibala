package utils

import (
	"auth/models"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func generateAid() string {
	id := uuid.New()
	return id.String()
}

func createToken(tokenType, Username, aid string) (string, error) {
	var JwtExpireTime int64
	switch tokenType {
	case "access":
		JwtExpireTime, _ = strconv.ParseInt(os.Getenv("JWT_ACCESS_EXPIRE_TIME"), 10, 32)
	case "refresh":
		JwtExpireTime, _ = strconv.ParseInt(os.Getenv("JWT_REFRESH_EXPIRE_TIME"), 10, 32)

	}

	claims := &models.Auth{
		Username:  Username,
		Aid:       aid,
		TokenType: fmt.Sprintf("%s_type", tokenType),

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(JwtExpireTime))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	createdToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return createdToken, err
}

func CreateTokens(userId string) (string, string, error) {
	generatedAid := generateAid()
	var err error
	var accessToken, refreshToken string
	accessToken, err = createToken("access", userId, generatedAid)
	refreshToken, err = createToken("refresh", userId, generatedAid)
	return accessToken, refreshToken, err

}

func GetTokenClaims(token string) (*models.Auth, error) {
	claims := models.Auth{}
	bearerToken := strings.Replace(token, "Bearer ", "", 1)
	_, err := jwt.ParseWithClaims(bearerToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, errors.New("token expired")
	}
	return &claims, nil

}

func IsValidToken(token string) (bool, *models.Auth, error) {
	claims, err := GetTokenClaims(token)
	if err != nil {
		return false, &models.Auth{}, errors.New("error in decode token")
	}
	return true, claims, nil
}
