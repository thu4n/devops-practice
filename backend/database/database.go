package database

import (
	"fmt"
	"go-backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	postgresdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the elephant. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to the elephant")
	postgresdb.Logger = logger.Default.LogMode(logger.Info) // Set log from db to info mode

	log.Println("Migrating data models")
	postgresdb.AutoMigrate(&models.User{}) // Migrate only User for now, more later

	DB = DbInstance{
		Db: postgresdb,
	}
}
