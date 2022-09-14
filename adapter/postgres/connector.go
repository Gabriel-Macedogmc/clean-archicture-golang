package postgres

import (
	"log"

	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func ConnectToDatabase(dsn string) *gorm.DB {
	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connected to Database...")

	return Instance
}

func Migrate() {
	Instance.AutoMigrate(&domain.Product{})
	log.Println("Database Migration Completed...")
}
