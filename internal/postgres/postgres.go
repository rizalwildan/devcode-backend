package postgres

import (
	"fmt"
	"github.com/rizalwildan/devcode-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func Connect() *gorm.DB {
	var err error
	p := config.GetEnv("PGSQL_PORT")
	port, err := strconv.Atoi(p)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("PGSQL_HOST"),
		port,
		config.GetEnv("PGSQL_USER"),
		config.GetEnv("PGSQL_PASSWORD"),
		config.GetEnv("PGSQL_DBNAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Duration(60) * time.Second)

	return db
}
