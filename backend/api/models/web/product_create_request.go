package web

type ProductCreate struct {
	Name          string  `json:"name"`
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price"`
	Stock         int     `json:"stock"`
}
