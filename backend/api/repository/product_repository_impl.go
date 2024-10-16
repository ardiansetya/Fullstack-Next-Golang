package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"jamal/api/api/models/domain"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (repository ProductRepositoryImpl) Create(tx *gorm.DB, product *domain.Product) (domain.Product, error) {
	result := tx.Create(&product)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}
	return *product, nil
}

func (repository ProductRepositoryImpl) Delete(tx *gorm.DB, productId int) error {
	// usage soft delete, and locking
	result := tx.Omit(clause.Associations).Where("id =?", productId).Delete(&domain.Product{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository ProductRepositoryImpl) Update(tx *gorm.DB, product *domain.Product, productId int) (domain.Product, error) {

	result := tx.Omit(clause.Associations).Where("id =? ", productId).Updates(&product)

	if result.Error != nil {
		return domain.Product{}, result.Error
	}

	if result.RowsAffected == 0 {
		return domain.Product{}, errors.New("no product found to update")
	}

	return *product, nil
}

func (repository ProductRepositoryImpl) FindById(tx *gorm.DB, productId int) (domain.Product, error) {
	var product domain.Product
	result := tx.Omit(clause.Associations).Take(&product, "id=?", productId)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}

	return product, nil
}

func (repository ProductRepositoryImpl) FindAll(tx *gorm.DB) ([]domain.Product, error) {
	var products []domain.Product
	result := tx.Find(&products)
	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}
