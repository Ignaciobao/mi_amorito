package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	GinMode        string
	MongoURI       string
	MongoDatabase  string
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
	OpenAIAPIKey   string
	OpenAIModel    string
	JWTSecret      string
	AllowedOrigins string
}

func Load() *Config {
	// 加载 .env 文件
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Port:           getEnv("PORT", "8080"),
		GinMode:        getEnv("GIN_MODE", "debug"),
		MongoURI:       getEnv("MONGODB_URI", "mongodb://admin:password123@localhost:27017/mi_amorito?authSource=admin"),
		MongoDatabase:  getEnv("MONGODB_DATABASE", "mi_amorito"),
		RedisAddr:      getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:  getEnv("REDIS_PASSWORD", ""),
		RedisDB:        0, // Redis DB 默认为 0
		OpenAIAPIKey:   getEnv("OPENAI_API_KEY", ""),
		OpenAIModel:    getEnv("OPENAI_MODEL", "gpt-4-1106-preview"),
		JWTSecret:      getEnv("JWT_SECRET", "default_jwt_secret"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:5173"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}