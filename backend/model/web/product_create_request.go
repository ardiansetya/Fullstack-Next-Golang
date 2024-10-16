package web

/// repesentation data yang akan direquest client (request body json)
// id auto inc

type ProductCreateRequest struct {
	Name        string  `validate:"required,max=200,min=1"`
	Description string  `validate:"max=200"`
	Price       float64 `validate:"required,min=1"`
	Category    string  `validate:"required,max=200"`
	Stock       int     `validate:"max=999"`
}
