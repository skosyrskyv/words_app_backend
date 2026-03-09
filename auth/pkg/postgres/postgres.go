package postgres

import (
	"auth/config"
	"fmt"
	"log"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func Init(cfg config.PostgresConfig) *Postgres {
	postgres, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})

	if err != nil {
		log.Fatal("DataBase Connection failed...")
	}

	fmt.Printf("Database connected")

	return &Postgres{
		db: postgres,
	}
}

func (p *Postgres) Close() {
	sqlDB, err := p.db.DB()
	if err != nil {
		slog.Error("Error getting sqlDB")
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
