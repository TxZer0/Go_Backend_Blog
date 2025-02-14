package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateAccessToken(userId uint) (string, error) {
	accessClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	return accessToken.SignedString(secretKey)
}

func GenerateRefreshToken(userId uint) (string, error) {
	refreshClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	return refreshToken.SignedString(secretKey)
}

func GenerateTokenPair(userId uint) (string, string, error) {
	accessToken, err := GenerateAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}

func GenerateEmailToken(email string, exp int64) (string, error) {
	accessClaims := jwt.MapClaims{
		"email": email,
		"exp":   exp,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	return accessToken.SignedString(secretKey)
}
