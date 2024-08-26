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

func Logout(c *gin.Context, db *gorm.DB) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	authUser, ok := user.(entities.AuthDAO)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// RefreshToken을 데이터베이스에서 제거
	if err := db.Model(&authUser).Update("refresh_token", "").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func DeleteAccount(c *gin.Context, db *gorm.DB) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	authUser, ok := user.(entities.AuthDAO)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// 트랜잭션 시작
	tx := db.Begin()

	// 사용자 관련 데이터 삭제
	if err := tx.Delete(&authUser).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user account"})
		return
	}

	// 여기에 사용자와 관련된 다른 데이터 삭제 로직 추가
	// 예: 사용자의 게시글, 댓글 등 삭제
	// if err := tx.Where("user_id = ?", authUser.UUID).Delete(&UserPosts{}).Error; err != nil {
	//     tx.Rollback()
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user data"})
	//     return
	// }

	// 트랜잭션 커밋
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account successfully deleted"})
}

func UpdateNickname(c *gin.Context, db *gorm.DB) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	authUser, ok := user.(entities.AuthDAO)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	var nicknameUpdate entities.NicknameUpdate
	if err := c.ShouldBindJSON(&nicknameUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 닉네임 업데이트
	if err := db.Model(&authUser).Update("nickname", nicknameUpdate.Nickname).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update nickname"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nickname successfully updated", "new_nickname": nicknameUpdate.Nickname})
}
