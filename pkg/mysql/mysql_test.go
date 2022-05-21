package mysql

import (
	"log"
	"testing"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestOpenDB(t *testing.T) {
	host, port, user, pass, dbName := "172.17.0.1", 3306, "root", "123456", "rec_movie"
	db, err := OpenDB(host, port, user, pass, dbName)
	if err != nil {
		log.Printf("OpenDB Failed: %v\n", err)
		return
	}
	if err = db.AutoMigrate(&Product{}); err != nil {
		log.Printf("AutoMigrate Failed: %v\n", err)
		return
	}
	// Create
	if err = db.Create(&Product{Code: "D42", Price: 100}).Error; err != nil {
		log.Printf("Create Failed: %v\n", err)
		return
	}
	var product Product
	if err = db.First(&product, "code = ?", "D42").Error; err != nil {
		log.Printf("Find Failed: %v\n", err)
		return
	}
	log.Printf("product = %+v\n", product)
}
