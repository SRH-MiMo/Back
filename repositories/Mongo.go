package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongo(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// MongoDB 클라이언트 생성
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// MongoDB에 연결
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// 연결 확인
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("MongoDB에 성공적으로 연결되었습니다!")

	return client, nil
}
