package postgres

import (
	"auth/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func MustOpenConnection(cfg config.AuthDBServerConfig) *gorm.DB {

	var err error

	Instance, err = gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})

	if err != nil {
		log.Fatal("DataBase Connection failed...")
	}

	fmt.Printf("Database connected")

	return Instance
}
