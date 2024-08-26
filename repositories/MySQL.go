package repositories

import (
	"Back/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entities.AuthDAO{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
