package service

import (
	"errors"

	"commercium/internal/entity"
	"commercium/internal/repository"
)

type ProductsService interface {
	GetAllProducts() ([]entity.Products, error)
	GetProductByID(id int) (entity.Products, error)
	GetProductByName(value string) ([]entity.Products, error)
	CreateProduct(product entity.ProductsRequest) (entity.Products, error)
	UpdateProduct(id int, productNew entity.ProductsRequest) (entity.Products, error)
	DeleteProduct(id int) (entity.Products, error)
}

type productsService struct {
	productsRepository repository.ProductsRepository
}

func NewProductsService(productsRepository repository.ProductsRepository) *productsService {
	return &productsService{productsRepository}
}

func (product_serv *productsService) GetAllProducts() ([]entity.Products, error) {
	return product_serv.productsRepository.GetAllProducts()
}

func (product_serv *productsService) GetProductByID(id int) (entity.Products, error) {
	return product_serv.productsRepository.GetProductByID(id)
}

func (product_serv *productsService) GetProductByName(value string) ([]entity.Products, error) {
	if value == "" {
		return nil, errors.New("parameter cannot be blank")
	}

	return product_serv.productsRepository.GetProductByName(value)
}

func (product_serv *productsService) CreateProduct(productRequest entity.ProductsRequest) (entity.Products, error) {
	if *productRequest.Stock < 0 {
		return entity.Products{}, errors.New("minimum product stock is 0")
	}

	product := entity.Products{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Stock:       *productRequest.Stock,
	}

	return product_serv.productsRepository.CreateProduct(product)
}

func (product_serv *productsService) UpdateProduct(id int, productNew entity.ProductsRequest) (entity.Products, error) {
	product, err := product_serv.productsRepository.GetProductByID(id)
	if err != nil {
		return entity.Products{}, errors.New("product not found")
	}

	// VALIDASI APAKAH ATTRIBUT PRODUCT SUDAH DI INPUT
	if productNew.Name != "" {
		product.Name = productNew.Name
	}
	if productNew.Description != "" {
		product.Description = productNew.Description
	}
	if productNew.Price != 0 {
		product.Price = productNew.Price
	}
	if productNew.Stock != nil {
		product.Stock = *productNew.Stock
	}

	return product_serv.productsRepository.UpdateProduct(product)
}

func (product_serv *productsService) DeleteProduct(id int) (entity.Products, error) {
	product, err := product_serv.productsRepository.GetProductByID(id)
	if err != nil {
		return entity.Products{}, errors.New("product not found")
	}

	return product_serv.productsRepository.DeleteProduct(product)
}
