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
}
