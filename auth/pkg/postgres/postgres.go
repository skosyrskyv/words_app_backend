package postgres

import (
	"auth/config"
	"auth/internal/infrastructure/postgres/models"
	"fmt"
	"log"
	"log/slog"

	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func Init(cfg config.PostgresConfig) *Postgres {
	postgres, err := gorm.Open(driver.Open(cfg.GetDSN()), &gorm.Config{})

	if err != nil {
		log.Fatal("DataBase Connection failed...")
	}

	fmt.Println("Database connected")

	if err := postgres.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	fmt.Println("Migrations completed")

	return &Postgres{
		db: postgres,
	}
}

func (p *Postgres) Close() {
	sqlDB, err := p.db.DB()
	if err != nil {
		slog.Error("Error getting sqlDB")
		return
	}

	err = sqlDB.Close()

	if err != nil {
		slog.Error("Error close Postgres DB")
	} else {
		slog.Info("Postgres DB connection closed")
	}
}

func (p *Postgres) DB() *gorm.DB {
	return p.db
}
