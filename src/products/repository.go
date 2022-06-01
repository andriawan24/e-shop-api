package products

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	GetProducts() ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProducts() ([]Product, error) {
	var products []Product

	err := r.db.Preload("ProductCategories.Category").Preload(clause.Associations).Find(&products).Error
	if err != nil {
		log.Fatal(err)
		return products, err
	}

	log.Println(products)
	return products, nil
}
