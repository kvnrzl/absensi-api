package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	DB_USERNAME  string
	DB_PASSWORD  string
	DB_HOST      string
	DB_PORT      string
	DB_NAME      string
	SECRET_KEY   string
	EXP_DURATION time.Duration
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	SECRET_KEY = os.Getenv("SECRET_KEY")
	duration, _ := strconv.Atoi(os.Getenv("EXP_DURATION"))
	EXP_DURATION = time.Hour * time.Duration(duration)
}
