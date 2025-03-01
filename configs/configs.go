package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	App App
	DB  DB
	S3  S3
}

type App struct {
	Port string
}

type DB struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

type S3 struct {
	CDNUrl       string
	BucketName   string
	BucketRegion string
	AccessKey    string
	SecretKey    string
}

func LoadConfig() *Configs {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Configs{
		App: App{
			Port: getEnv("PORT", "3333"),
		},
		DB: DB{
			DBUser:     getEnv("DB_USER", ""),
			DBPassword: getEnv("DB_PASSWORD", ""),
			DBHost:     getEnv("DB_HOST", ""),
			DBPort:     getEnv("DB_PORT", ""),
			DBName:     getEnv("DB_NAME", ""),
		},
		S3: S3{
			CDNUrl:       getEnv("S3_CDN_URL", ""),
			BucketName:   getEnv("S3_BUCKET_NAME", ""),
			BucketRegion: getEnv("S3_BUCKET_REGION", ""),
			AccessKey:    getEnv("S3_ACCESS_KEY_ID", ""),
			SecretKey:    getEnv("S3_SECRET_ACCESS_KEY", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
