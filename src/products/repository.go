package products

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	GetProducts(categoryId int) ([]Product, error)
	GetCategories() ([]Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProducts(categoryId int) ([]Product, error) {
	var products []Product

	db := r.db.Preload("Category")

	if categoryId > 0 {
		db = db.Where("category_id = ?", categoryId)
	}

	err := db.Preload(clause.Associations).Find(&products).Error
	if err != nil {
		log.Fatal(err)
		return products, err
	}

	return products, nil
}

func (r *repository) GetCategories() ([]Category, error) {
	var categories []Category

	err := r.db.Find(&categories).Error
	if err != nil {
		log.Fatal(err)
		return categories, err
	}

	return categories, nil
}
