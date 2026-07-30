package configs

import (
	"log"

	"os"

	"github.com/joho/godotenv"
)


type ConfigStruct struct {

	DBHost string

	DBUser string

	DBPassword string

	DBName string

	DBPort string
}


var Config ConfigStruct


func LoadConfig() {


	err := godotenv.Load()

	if err != nil {

		log.Println("No .env file found")

	}


	Config = ConfigStruct{

		DBHost: os.Getenv("DB_HOST"),

		DBUser: os.Getenv("DB_USER"),

		DBPassword: os.Getenv("DB_PASSWORD"),

		DBName: os.Getenv("DB_NAME"),

		DBPort: os.Getenv("DB_PORT"),

	}

}