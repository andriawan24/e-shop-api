package carts

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	GetUserCart(userID int) (Cart, error)
	CreateOrGetCartByUserID(userID int) (Cart, error)
	SaveCart(detail CartDetail, cart Cart) (Cart, error)
	RemoveCart(cart Cart) (bool, error)
	RemoveProduct(cartId int, projectId int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUserCart(userID int) (Cart, error) {
	var cart Cart

	err := r.db.Preload("CartDetail", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Preload(clause.Associations)
		})
	}).Preload(clause.Associations).Find(&cart).Error
	if err != nil {
		return cart, err
	}

	if cart.ID == 0 {
		return cart, errors.New("Cart not found")
	}

	return cart, nil
}

func (r *repository) CreateOrGetCartByUserID(userID int) (Cart, error) {
	var cart Cart

	err := r.db.Where("user_id = ?", userID).Find(&cart).Error
	if err != nil {
		return cart, err
	}

	if cart.ID == 0 {
		cart = Cart{
			UserID: userID,
		}
		err = r.db.Save(&cart).Error
		if err != nil {
			return cart, err
		}
	}

	return cart, nil
}

func (r *repository) SaveCart(detail CartDetail, cart Cart) (Cart, error) {
	var findByIdProduct CartDetail
	err := r.db.Where("product_id = ?", detail.ProductID).Find(&findByIdProduct).Error
	if err != nil {
		return cart, err
	}

	if findByIdProduct.ID > 0 {
		detail.ID = findByIdProduct.ID
	}

	err = r.db.Save(&detail).Error
	if err != nil {
		return cart, err
	}

	err = r.db.Preload("CartDetail", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Preload(clause.Associations)
		})
	}).Preload(clause.Associations).Find(&cart).Error

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *repository) RemoveCart(cart Cart) (bool, error) {

	err := r.db.Delete(&cart).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) RemoveProduct(cartId int, projectId int) (bool, error) {

	var detail CartDetail
	err := r.db.Where("cart_id = ?", cartId).Where("product_id = ?", projectId).Find(&detail).Error
	if err != nil {
		return false, err
	}

	err = r.db.Delete(&detail).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
