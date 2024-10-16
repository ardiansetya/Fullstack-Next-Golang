package repository

import (
	"gorm.io/gorm"
	"jamal/api/api/models/domain"
)

type ProductRepository interface {
	Create(tx *gorm.DB, product *domain.Product) (domain.Product, error)
	Delete(tx *gorm.DB, productId int) error
	Update(tx *gorm.DB, product *domain.Product, productId int) (domain.Product, error)
	FindById(tx *gorm.DB, productId int) (domain.Product, error)
	FindAll(tx *gorm.DB) ([]domain.Product, error)
}
