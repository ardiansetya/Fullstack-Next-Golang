package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"korie/api-product/exception"
	"korie/api-product/helper"
	"korie/api-product/model/domain"
	"korie/api-product/model/web"
	"korie/api-product/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository //<-- butuh untuk manipulasi data
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository, DB: DB, Validate: validate}
}

/// INGAT KITA MENGUNAKAN TX  ///

func (service ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	//<-- validate err
	err := service.Validate.Struct(request)
	helper.IfErrNotNul(err)

	tx, err := service.DB.Begin() //<-- start tx
	helper.IfErrNotNul(err)
	defer helper.CommitOrRollback(tx, err)

	// <-- menampung data request user
	product := domain.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Category:    request.Category,
		Stock:       request.Stock,
	}

	//<-- execute query
	product = service.ProductRepository.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	//<-- validate err
	err := service.Validate.Struct(request)
	helper.IfErrNotNul(err)

	tx, err := service.DB.Begin() //<-- start tx
	helper.IfErrNotNul(err)
	defer helper.CommitOrRollback(tx, err)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// <-- menampung data request user
	product = domain.Product{
		Id:          request.Id,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Category:    request.Category,
		Stock:       request.Stock,
	}

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin() //<-- start tx
	helper.IfErrNotNul(err)
	defer helper.CommitOrRollback(tx, err)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product)

}

func (service ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin() //<-- start tx
	helper.IfErrNotNul(err)
	defer helper.CommitOrRollback(tx, err)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin() //<-- start tx
	helper.IfErrNotNul(err)
	defer helper.CommitOrRollback(tx, err)

	products := service.ProductRepository.FindAll(ctx, tx)

	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, helper.ToProductResponse(product))
	}
	return productResponses
}
