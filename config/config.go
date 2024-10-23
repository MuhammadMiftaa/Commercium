package config

import (
	"commercium/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=123 dbname=commercium port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&entity.Users{}, &entity.Products{}, &entity.Orders{})

	return db, err
}

func Migrate(db *gorm.DB, entity ...interface{}) {
    db.AutoMigrate(entity...)
}