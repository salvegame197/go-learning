package services

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DOTenv interface {
	Get(key string) string
	hasEnv(key string) bool
}

func DOTenvService() DOTenv {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func (d *dotenv) Get(key string) string {
	return os.Getenv(key)
}
func (d *dotenv) hasEnv(key string) bool {
	return os.Getenv(key) != ""
}
