package web

type ProductUpdateRequest struct {
	Id          int     `validate:"required"`
	Name        string  `validate:"required,max=200,min=1"`
	Description string  `validate:"max=200"`
	Price       float64 `validate:"required,min=1"`
	Category    string  `validate:"required,max=200"`
	Stock       int     `validate:"max=999"`
}
