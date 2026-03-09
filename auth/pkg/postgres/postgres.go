package postgres

import (
	"auth/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg config.PostgresConfig) *gorm.DB {
	postgres, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})

	if err != nil {
		log.Fatal("DataBase Connection failed...")
	}

	fmt.Printf("Database connected")

	return postgres
}
