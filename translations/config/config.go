package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	LOCAL = "local"
	PROD  = "prod"
	DEV   = "dev"
)

type Config struct {
	Env              string `yaml:"env" env-default:"local" env-required:"true"`
	HTTPClientConfig `yaml:"http_client"`
	GRPCServerConfig `yaml:"grpc_server"`
	PostgresConfig   `yaml:"postgres"`
	JWTConfig        `yaml:"jwt"`
	RedisConfig      `yaml:"redis"`
}

type HTTPClientConfig struct {
	ReadTimeout int `yaml:"read_timeout" env:"HTTP_CLIENT_READ_TIMEOUT" env-default:"5"`
}

type GRPCServerConfig struct {
	Address string `yaml:"address" env:"GRPC_HOST"`
	Port    string `yaml:"port"    env:"GRPC_PORT"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"     env:"POSTGRES_HOST"     env-default:"localhost"`
	Port     string `yaml:"port"     env:"POSTGRES_PORT"     env-default:"5432"`
	Name     string `yaml:"name"     env:"POSTGRES_DB"`
	User     string `yaml:"user"     env:"POSTGRES_USER"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
	Sslmode  string `yaml:"sslmode"`
	Timezone string `yaml:"timezone"`
}

type JWTConfig struct {
	PrivateKeyPath string `yaml:"private_key_path"`
	PublicKeyPath  string `yaml:"public_key_path"`
	Issuer         string `env:"JWT_ISSUER"`
	AccessTTL      string `yaml:"access_ttl"`
	RefreshTTL     string `yaml:"refresh_ttl"`
}

type RedisConfig struct {
	Host         string `yaml:"host"     env:"REDIS_HOST"`
	Port         string `yaml:"port"     env:"REDIS_PORT"`
	Password     string `yaml:"password" env:"REDIS_PASSWORD"`
	DB           int    `yaml:"db"       env:"REDIS_DB"`
	Username     string `yaml:"username"`
	MaxRetries   int    `yaml:"max_retries"`
	DialTimeout  int    `yaml:"dial_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

func Init() Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using environment variables: %s", err)
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

func (c *JWTConfig) GetAccessTTL() (time.Duration, error) {
	duration, err := c.parseDuration(c.AccessTTL)
	if err != nil {
		return 15 * time.Minute, err
	}
	return duration, nil
}

func (c *JWTConfig) GetRefreshTTL() (time.Duration, error) {
	duration, err := c.parseDuration(c.RefreshTTL)
	if err != nil {
		return 30 * 24 * time.Hour, err
	}
	return duration, nil
}

func (c *JWTConfig) parseDuration(s string) (time.Duration, error) {
	if len(s) > 1 && s[len(s)-1] == 'd' {
		days := 0
		n, err := fmt.Sscanf(s, "%dd", &days)
		if err != nil || n != 1 || days <= 0 {
			return 0, fmt.Errorf("invalid duration: %s", s)
		}
		return time.Duration(days) * 24 * time.Hour, nil
	}
	return time.ParseDuration(s)
}
