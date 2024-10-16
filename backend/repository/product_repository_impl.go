package repository

import (
	"context"
	"database/sql"
	"errors"
	"korie/api-product/helper"
	"korie/api-product/model/domain"
)

/// IMPLEMENT METHOD YANG DI BUAT DI product_repository.go ///

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products (name, description, price, category, stock) values (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Description, product.Price, product.Category, product.Stock)
	helper.IfErrNotNul(err)

	id, err := result.LastInsertId()
	helper.IfErrNotNul(err)

	product.Id = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE products SET name=?, description=?, price=?, category=?, stock=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Description, product.Price, product.Category, product.Stock, product.Id)
	helper.IfErrNotNul(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "DELETE FROM products WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.IfErrNotNul(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "SELECT * FROM products WHERE id =?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.IfErrNotNul(err)
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Category, &product.Stock)
		helper.IfErrNotNul(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT * FROM products"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.IfErrNotNul(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Category, &product.Stock)
		helper.IfErrNotNul(err)
		products = append(products, product)
	}

	return products
}
