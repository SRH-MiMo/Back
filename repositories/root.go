package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func NewRepository(dns string, uri string) (*gorm.DB, *mongo.Client, error) {
	sql, err := NewMySQL(dns)
	if err != nil {
		return nil, nil, err
	}

	mmong, err := NewMongo(uri)
	if err != nil {
		return nil, nil, err
	}

	return sql, mmong, nil
}
