package services

import (
	"Back/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func KakaoLogin(c *gin.Context, db *gorm.DB) {
	var user *entities.AuthDTO
	var dbUser *entities.AuthDAO
	var status string

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 이메일로 사용자 검색
	result := db.Where("nickname = ?", user.Nickname).First(&dbUser)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// 사용자가 없으면 새로 생성
			newUser := &entities.AuthDAO{
				UUID:     uuid.New().String(),
				Nickname: user.Nickname,
				Age:      user.Age,
				Job:      user.Job,
			}

			if err := db.Create(newUser).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			dbUser = newUser
			status = "유저생성됨"
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	} else {
		// 사용자가 이미 존재함
		status = "있던 유저임"
	}

	// JWT 토큰 생성
	tokenPair, err := generateTokenPair(dbUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Refresh 토큰을 DB에 저장
	dbUser.RefreshToken = tokenPair.RefreshToken
	if err := db.Save(dbUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save refresh token"})
		return
	}

	// 상태 정보를 포함한 응답 생성
	response := gin.H{
		"status": status,
		"token":  tokenPair,
	}

	c.JSON(http.StatusOK, response)
}
