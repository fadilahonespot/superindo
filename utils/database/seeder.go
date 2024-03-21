package database

import (
	"log"
	"os"

	"github.com/fadilahonespot/superindo/entity"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) *gorm.DB {
	var categoryCount int64
	err := db.Model(&entity.Category{}).Count(&categoryCount).Error
	if err != nil {
		log.Fatal(err)
	}
	if categoryCount == 0 {
		sqlFile, err := os.ReadFile("./resources/categories.sql")
		if err != nil {
			log.Fatal(err)
		}
		err = db.Exec(string(sqlFile)).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	var productCount int64
	err = db.Model(&entity.Product{}).Count(&productCount).Error
	if err != nil {
		log.Fatal(err)
	}
	if productCount == 0 {
		sqlFile, err := os.ReadFile("./resources/products.sql")
		if err != nil {
			log.Fatal(err)
		}
		err = db.Exec(string(sqlFile)).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}
