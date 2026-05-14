package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host              string
	Port              string
	AdminPassword     string
	API_KEY           string
	API_SECRET        string
	CLOUD_NAME        string
	CLOUDINARY_FOLDER string
}

func GetEnv() *Env {
	return &Env{
		Host:              os.Getenv("HOST"),
		Port:              os.Getenv("PORT"),
		AdminPassword:     os.Getenv("ADMIN_PASSWORD"),
		API_KEY:           os.Getenv("API_KEY"),
		API_SECRET:        os.Getenv("API_SECRET"),
		CLOUD_NAME:        os.Getenv("CLOUD_NAME"),
		CLOUDINARY_FOLDER: os.Getenv("CLOUDINARY_FOLDER"),
	}
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .Env file.")
	}
}
