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

	JWTSecret string

	HospitalAAPIURL string
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

		JWTSecret: os.Getenv("JWT_SECRET"),

		HospitalAAPIURL: os.Getenv("HOSPITAL_A_API_URL"),
	}

}
