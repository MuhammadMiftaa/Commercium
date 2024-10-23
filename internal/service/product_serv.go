package service

import (
	"errors"
	"strconv"

	"commercium/internal/entity"
	"commercium/internal/repository"
)

type ProductsService interface {
	GetAllProducts() ([]entity.Products, error)
	GetProductByID(id int) (entity.Products, error)
	CreateProduct(product entity.Products) (entity.Products, error)
	UpdateProduct(product entity.Products) (entity.Products, error)
	DeleteProduct(product entity.Products) (entity.Products, error)
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

func (product_serv *productsService) CreateProduct(product entity.Products) (entity.Products, error) {
	return product_serv.productsRepository.CreateProduct(product)
}

func (product_serv *productsService) UpdateProduct(id int, productNew entity.Products) (entity.Products, error) {
	product, err := product_serv.productsRepository.GetProductByID(id)
	if err != nil {
		return productNew, errors.New("Product not found")
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
	if strconv.Itoa(productNew.Stock) != "" {
		product.Stock = productNew.Stock
	}

	return product_serv.productsRepository.UpdateProduct(product)
}

func (product_serv *productsService) DeleteProduct(id int) (entity.Products, error) {
	product, err := product_serv.productsRepository.GetProductByID(id)
	if err != nil {
		return product, errors.New("Product not found")
	}

	return product_serv.productsRepository.DeleteProduct(product)
}
