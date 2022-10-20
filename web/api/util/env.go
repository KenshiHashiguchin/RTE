package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Env(param string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	result := os.Getenv(param)
	return result
}
