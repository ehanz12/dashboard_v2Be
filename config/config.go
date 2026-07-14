package config

import (
	"os"
	"strconv"
	"github.com/joho/godotenv"
)


type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	Port	   string
	RedisHost  string
	RedisPort  string
	RedisPassword string
	RedisDB    int
}

var AppConfig *Config

func LoadEnv() {
	//load env
	if err := godotenv.Load(); err != nil  {
		panic("Failed to load .env file")
	} 
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	AppConfig = &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		Port:       os.Getenv("PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB: db,
	}
}
