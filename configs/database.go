package configs

import (
	"fmt"
	"log"

	model "github.com/rungrawee-popcorn/hospital-middleware-system/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB


func ConnectDatabase() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Config.DBHost,
		Config.DBUser,
		Config.DBPassword,
		Config.DBName,
		Config.DBPort,
	)


	database, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)


	if err != nil {

		log.Fatal("Failed to connect database:", err)

	}


	DB = database


	log.Println("Database connected successfully")
}



func AutoMigrate() {

	err := DB.AutoMigrate(
		&model.Hospital{},
		&model.Staff{},
		&model.Patient{},
	)


	if err != nil {

		log.Fatal("Migration failed:", err)

	}


	log.Println("Migration completed successfully")
}