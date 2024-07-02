package mysql

import (
	"fmt"
	"github.com/rizalwildan/devcode-backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func Connect() *gorm.DB {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("MYSQL_USER"),
		config.GetEnv("MYSQL_PASSWORD"),
		config.GetEnv("MYSQL_HOST"),
		config.GetEnv("MYSQL_PORT"),
		config.GetEnv("MYSQL_DBNAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Duration(60) * time.Second)

	return db
}
