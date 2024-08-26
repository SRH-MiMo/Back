package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func NewRepository(dns string, uri string) (*gorm.DB, *mongo.Client) {
	sql, err := NewMySQL(dns)
	if err != nil {
		panic(err)
	}

	mmong, err := NewMongo(uri)
	if err != nil {
		panic(err)
	}

	return sql, mmong
}
