package database

import (
	"log"
	"os"

	"github.com/fadilahonespot/superindo/entity"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) *gorm.DB {
	var categoryCount int64
	tx := db.Begin()
	err := db.Model(&entity.Category{}).Count(&categoryCount).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	if categoryCount == 0 {
		sqlFile, err := os.ReadFile("./resources/categories.sql")
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		tx = tx.Exec(string(sqlFile))
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	var productCount int64
	err = db.Model(&entity.Product{}).Count(&productCount).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	if productCount == 0 {
		sqlFile, err := os.ReadFile("./resources/products.sql")
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		tx = tx.Exec(string(sqlFile))
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	tx.Commit()
	return tx
}
