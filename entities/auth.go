package entities

import (
	"errors"
	"os"
	"time"
)

type AuthDAO struct {
	UUID         string `gorm:"primary_key"`
	Nickname     string `gorm:"primary_key"`
	Age          string
	Job          string
	RefreshToken string
}

type AuthDTO struct {
	Nickname string `json:"nickname"`
	Age      string `json:"age"`
	Job      string `json:"job"`
}

var (
	JwtSecret       = []byte(os.Getenv("JWT_SECRET"))
	AccessTokenExp  = time.Hour
	RefreshTokenExp = time.Hour * 24 * 7
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type NicknameUpdate struct {
	Nickname string `json:"nickname" binding:"required"`
}
