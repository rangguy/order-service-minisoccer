package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
	"time"
)

func InitDatabase() (*gorm.DB, error) {
	config := Config
	encodedPassword := url.QueryEscape(config.Database.Password)
	uri := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Username,
		encodedPassword,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConnection)
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Database.MaxLifetimeConnection) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.Database.MaxIdleConnection) * time.Second)

	return db, nil
}
