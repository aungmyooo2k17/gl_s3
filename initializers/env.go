package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func RemoteUrl() string {
	return GetEnvWithKey("SERVER") + ":" + GetEnvWithKey("PORT")
}
