package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	LOCAL = "local"
	PROD  = "prod"
	DEV   = "dev"
)

type Config struct {
	Env              string `yaml:"env" env: "ENV" env-default:"local" env-required:"true"`
	HTTPServerConfig `yaml:"http_server"`
	PostgresConfig   `yaml:"postgres"`
}

type HTTPServerConfig struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type PostgresConfig struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
	Timezone string `yaml:"timezone"`
}

func Init() Config {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is missing in environment variables")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file [%s] doesn't exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Error while reading config file: %s", err)
	}

	return cfg
}

func (c PostgresConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
		c.Host,
		c.Port,
		c.User,
		c.Name,
		c.Password,
		c.Sslmode,
		c.Timezone,
	)
}
