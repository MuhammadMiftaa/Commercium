package repository

import (
	"commercium/internal/entity"

	"gorm.io/gorm"
)

type ProductsRepository interface {
	GetAllProducts() ([]entity.Products, error)
	GetProductByID(id int) (entity.Products, error)
	CreateProduct(product entity.Products) (entity.Products, error)
	UpdateProduct(product entity.Products) (entity.Products, error)
	DeleteProduct(product entity.Products) (entity.Products, error)
}

type productsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *productsRepository {
	return &productsRepository{db}
}

func (product_repo *productsRepository) GetAllProducts() ([]entity.Products, error) {
	var products []entity.Products
	err := product_repo.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (product_repo *productsRepository) GetProductByID(id int) (entity.Products, error) {
	var product entity.Products
	err := product_repo.db.First(&product, id).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (product_repo *productsRepository) CreateProduct(product entity.Products) (entity.Products, error) {
	err := product_repo.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (product_repo *productsRepository) UpdateProduct(product entity.Products) (entity.Products, error) {
	err := product_repo.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (product_repo *productsRepository) DeleteProduct(product entity.Products) (entity.Products, error) {
	err := product_repo.db.Delete(&product).Error
	if err != nil {
		return product, nil
	}

	return product, nil
}
