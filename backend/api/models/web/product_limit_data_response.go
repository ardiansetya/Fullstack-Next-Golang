package web

import (
	"jamal/api/api/models/domain"
)

// MEMBATASI DATA YANG DI BAGIKAN KE USER

type ProductLimitData struct {
	Name          string  `json:"name"`
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price"`
	Stock         int     `json:"stock"`
}

// Fungsi untuk membuat ProductLimitData dari Product

func NewProductLimitData(product domain.Product) ProductLimitData {
	return ProductLimitData{
		Name:          product.Name,
		PurchasePrice: product.PurchasePrice,
		SellingPrice:  product.SellingPrice,
		Stock:         product.Stock,
	}
}
