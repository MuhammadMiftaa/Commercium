package repository

import (
	"commercium/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type ProductsRepository interface {
	GetAllProducts() ([]entity.Products, error)
	GetProductByID(id int) (entity.Products, error)
	GetProductByName(value string) ([]entity.Products, error)
	CreateProduct(product entity.Products) (entity.Products, error)
	UpdateProduct(product entity.Products) (entity.Products, error)
	DeleteProduct(product entity.Products) (entity.Products, error)
}

type productsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) ProductsRepository {
	return &productsRepository{db}
}

func (product_repo *productsRepository) GetAllProducts() ([]entity.Products, error) {
	var products []entity.Products
	err := product_repo.db.Find(&products).Error
	if err != nil {
		return nil, errors.New("failed to get products")
	}

	return products, nil
}

func (product_repo *productsRepository) GetProductByID(id int) (entity.Products, error) {
	var product entity.Products
	err := product_repo.db.First(&product, id).Error
	if err != nil {
		return entity.Products{}, errors.New("product not found")
	}

	return product, nil
}

func (product_repo *productsRepository) GetProductByName(value string) ([]entity.Products, error) {
	var products []entity.Products
	err := product_repo.db.Where("name LIKE ?", "%"+value+"%").Find(&products).Error
	if err != nil {
		return nil, errors.New("product not found")
	}

	return products, nil
}

func (product_repo *productsRepository) CreateProduct(product entity.Products) (entity.Products, error) {
	err := product_repo.db.Create(&product).Error
	if err != nil {
		return entity.Products{}, errors.New("failed to create product")
	}

	return product, nil
}

func (product_repo *productsRepository) UpdateProduct(product entity.Products) (entity.Products, error) {
	err := product_repo.db.Save(&product).Error
	if err != nil {
		return entity.Products{}, errors.New("failed to update product")
	}

	return product, nil
}

func (product_repo *productsRepository) DeleteProduct(product entity.Products) (entity.Products, error) {
	err := product_repo.db.Delete(&product).Error
	if err != nil {
		return entity.Products{}, errors.New("failed to delete product")
	}

	return product, nil
}
