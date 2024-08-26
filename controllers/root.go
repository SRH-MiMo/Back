package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func NewContorllers(port string) error {
	// 라우터 생성
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		MaxAge:       24 * time.Hour,
	}))

	v1 := r.Group("/api/v1")
	{

	}

	err := r.Run(port)
	return err
}
