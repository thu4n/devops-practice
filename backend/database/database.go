package database

import (
	"encoding/json"
	"fmt"
	"go-backend/models"
	"io"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Postgres DbInstance

func Init() {
	ConnectDb()
	PopulateDrinks("data/drinks.json")
}

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the elephant. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to the elephant")
	db.Logger = logger.Default.LogMode(logger.Info) // Set log from db to info mode

	log.Println("Migrating data models")
	db.AutoMigrate(
		&models.User{},
		&models.Drink{},
		&models.DrinkSize{},
		&models.Order{},
		&models.OrderItem{},
	)

	Postgres = DbInstance{
		Db: db,
	}

}

// Insert the sample JSON data into the elephant db instance
func PopulateDrinks(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("failed to open JSON file: %w", err)
		os.Exit(2)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var drinks []models.Drink

	err = json.Unmarshal(byteValue, &drinks)
	if err != nil {
		log.Fatal("failed to parse JSON: %w", err)
		os.Exit(2)
	}

	for _, drink := range drinks {
		if err := Postgres.Db.Create(&drink).Error; err != nil {
			log.Fatal("failed to insert drink: %w", err)
			os.Exit(2)
		}
	}

	log.Println("Database populated successfully")
}
