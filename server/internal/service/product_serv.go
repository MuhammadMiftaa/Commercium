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

func NewProductsService(productsRepository repository.ProductsRepository) ProductsService {
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
	// VALIDASI APAKAH PRODUCT NAME, DESCRIPTION, PRICE, STOCK KOSONG
	if productRequest.Name == "" || productRequest.Description == "" || productRequest.Price == 0 || productRequest.Stock == nil || productRequest.Category == "" {
		return entity.Products{}, errors.New("product name, description, price, stock, and category cannot be blank")
	}

	// VALIDASI APAKAH PRODUCT PRICE TIDAK NEGATIF
	if productRequest.Price < 0 {
		return entity.Products{}, errors.New("product price cannot be negative")
	}

	// VALIDASI APAKAH PRODUCT STOCK TIDAK NEGATIF
	if productRequest.Stock == nil || *productRequest.Stock < 0 {
		return entity.Products{}, errors.New("product stock cannot be negative")
	}

	product := entity.Products{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Stock:       *productRequest.Stock,
		Category:  productRequest.Category,
	}

	return product_serv.productsRepository.CreateProduct(product)
}

func (product_serv *productsService) UpdateProduct(id int, productNew entity.ProductsRequest) (entity.Products, error) {
	product, err := product_serv.productsRepository.GetProductByID(id)
	if err != nil {
		return entity.Products{}, errors.New("product not found")
	}

	// VALIDASI APAKAH PRODUCT NAME, DESCRIPTION, PRICE, STOCK KOSONG
	if productNew.Name == "" || productNew.Description == "" || productNew.Price == 0 || productNew.Stock == nil || productNew.Category == "" {
		return entity.Products{}, errors.New("product name, description, price, stock, and category cannot be blank")
	}

	// VALIDASI APAKAH PRODUCT PRICE TIDAK NEGATIF
	if productNew.Price < 0 {
		return entity.Products{}, errors.New("product price cannot be negative")
	}

	// VALIDASI APAKAH PRODUCT STOCK TIDAK NEGATIF
	if productNew.Stock == nil || *productNew.Stock < 0 {
		return entity.Products{}, errors.New("product stock cannot be negative")
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
	if productNew.Category != "" {
		product.Category = productNew.Category
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

