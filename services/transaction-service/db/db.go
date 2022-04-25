package db

import (
	"fmt"
	"log"

	"github.com/arvryna/betnomi/transaction-service/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Just need DB url to init the database and also perform necessary
// migrations
func Init() *gorm.DB {
	// TODO: move this to config
	uName := "betnomiadmin"
	pWord := "asd279364kk"
	dbURL := fmt.Sprintf("postgres://%s:%s@database:5432/betnomi", uName, pWord)

	// open connection:
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		// Use a better logging tool, eg: Logrus
		log.Println("DB is initialized!")
	}

	// WARN: remove this in prod
	db.AutoMigrate(&model.Transaction{})
	return db
}
