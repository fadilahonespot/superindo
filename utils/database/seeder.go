package database

import (
	"log"
	"os"
	"time"

	"github.com/fadilahonespot/superindo/entity"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) *gorm.DB {
	var categoryCount int64
	err := db.Model(&entity.Category{}).Count(&categoryCount).Error
	if err != nil {
		panic(err)
	}
	if categoryCount == 0 {
		sqlFile, err := os.ReadFile("./resources/categories.sql")
		if err != nil {
			log.Fatal(err)
		}
		db = db.Exec(string(sqlFile))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(4) * time.Second)
	}

	var productCount int64
	err = db.Model(&entity.Product{}).Count(&productCount).Error
	if err != nil {
		panic(err)
	}
	if productCount == 0 {
		sqlFile, err := os.ReadFile("./resources/products.sql")
		if err != nil {
			log.Fatal(err)
		}
		db = db.Exec(string(sqlFile))
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}
