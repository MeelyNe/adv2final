package config

import (
	"os"
	"sync"
)

type Config struct {
	DBUser  string
	DBPass  string
	DBHost  string
	DBPort  string
	DBName  string
	AppPort string
}

var (
	cfg  *Config
	once *sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{
			DBUser:  os.Getenv("DB_USER"),
			DBPass:  os.Getenv("DB_PASS"),
			DBHost:  os.Getenv("DB_HOST"),
			DBPort:  os.Getenv("DB_PORT"),
			DBName:  os.Getenv("DB_NAME"),
			AppPort: os.Getenv("APP_PORT"),
		}
	})

	return cfg
}
