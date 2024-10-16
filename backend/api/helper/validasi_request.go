package helper

import (
	web2 "jamal/api/api/models/web"
	"net/http"
)

// product create ama update filed sama makanya buat 1 saja

// Fungsi untuk memvalidasi data input produk
func validateProductRequest(productReq web2.ProductCreate) *web2.WebResponse {
	if productReq.Name == "" {
		return &web2.WebResponse{Code: http.StatusBadRequest, Status: "BAD REQUEST", Data: nil}
	}
	if productReq.PurchasePrice < 0 {
		return &web2.WebResponse{Code: http.StatusBadRequest, Status: "BAD REQUEST", Data: nil}
	}
	if productReq.SellingPrice < 0 {
		return &web2.WebResponse{Code: http.StatusBadRequest, Status: "BAD REQUEST", Data: nil}
	}

	return nil // Tidak ada kesalahan, validasi berhasil
}
