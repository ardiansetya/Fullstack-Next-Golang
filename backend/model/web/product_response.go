package web

/// ProductResponse ///
// -> di buat agar bisa membatasi data apa yang ingin diexpose ke client

type ProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
}
