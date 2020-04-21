package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_ENV        string
	PORT           string
	JWT_KEY        string
	JWT_EXP        string
	MYSQL_URL      string
	DB_HOST        string
	DB_NAME        string
	DB_USER        string
	DB_PASSWORD    string
	FRONTEND_URL   string
	EMAIL_SERVER   string
	EMAIL_PORT     string
	EMAIL_USER     string
	EMAIL_PASSWORD string
	DATA_DIR       string
)

func init() {
	godotenv.Load()

	APP_ENV = os.Getenv("APP_ENV")
	PORT = os.Getenv("PORT")
	JWT_KEY = os.Getenv("JWT_KEY")
	JWT_EXP = os.Getenv("JWT_EXP")
	MYSQL_URL = os.Getenv("MYSQL_URL")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	FRONTEND_URL = os.Getenv("FRONTEND_URL")
	EMAIL_SERVER = os.Getenv("EMAIL_SERVER")
	EMAIL_PORT = os.Getenv("EMAIL_PORT")
	EMAIL_USER = os.Getenv("EMAIL_USER")
	EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")
	DATA_DIR = os.Getenv("DATA_DIR")
}
