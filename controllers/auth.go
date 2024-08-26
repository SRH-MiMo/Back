package controllers

import (
	services "Back/serivces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func NewAuthController(r *gin.RouterGroup, sql *gorm.DB, nsql *mongo.Client) {
	r.POST("/login", func(c *gin.Context) {
		services.KakaoLogin(c, sql)
	})

	r.POST("/logout", func(c *gin.Context) {
		services.Logout(c, sql)
	})

	r.DELETE("/del", func(c *gin.Context) {
		services.DeleteAccount(c, sql)
	})

	r.PATCH("/update", func(c *gin.Context) {
		services.UpdateNickname(c, sql)
	})
}
