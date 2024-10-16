package repository

import (
	"context"
	"database/sql"
	"korie/api-product/model/domain"
)

//<-- buat contact terlebih dahulu
// kita mengunakan tx karena agar jika terjadi kegagalan query bisa rollback

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
