package serivces

import (
	"Back/entities"
	"github.com/golang-jwt/jwt"
	"time"
)

func generateTokenPair(user *entities.AuthDAO) (*entities.TokenPair, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":     user.UUID,
		"nickname": user.Nickname,
		"exp":      time.Now().Add(entities.AccessTokenExp).Unix(),
		"type":     "access",
	})

	accessTokenString, err := accessToken.SignedString(entities.JwtSecret)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":     user.UUID,
		"nickname": user.Nickname,
		"exp":      time.Now().Add(entities.AccessTokenExp).Unix(),
		"type":     "access",
	})

	refreshTokenString, err := refreshToken.SignedString(entities.JwtSecret)
	if err != nil {
		return nil, err
	}

	return &entities.TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
