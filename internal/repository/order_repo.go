package repository

import (
	"time"

	"commercium/internal/entity"

	"gorm.io/gorm"
)

type OrdersRepository interface {
	GetAllOrders() ([]entity.Orders, error)
	GetOrderByID(id int) (entity.Orders, error)
	GetOrderByDate(from, to time.Time) ([]entity.Orders, error)
	CreateOrder(order entity.Orders) (entity.Orders, error)
	UpdateOrder(order entity.Orders) (entity.Orders, error)
	DeleteOrder(order entity.Orders) (entity.Orders, error)
}

type ordersRepository struct {
	db *gorm.DB
}

func NewOrdersRepository(db *gorm.DB) *ordersRepository {
	return &ordersRepository{db}
}

func (order_repo *ordersRepository) GetAllOrders() ([]entity.Orders, error) {
	var orders []entity.Orders
	err := order_repo.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (order_repo *ordersRepository) GetOrderByID(id int) (entity.Orders, error) {
	var order entity.Orders
	err := order_repo.db.First(&order, id).Error
	if err != nil {
		return entity.Orders{}, err
	}

	return order, nil
}

func (order_repo *ordersRepository) GetOrderByDate(from, to time.Time) ([]entity.Orders, error) {
	var orders []entity.Orders
	err := order_repo.db.Where("created_at BETWEEN ? AND ?", from, to).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (order_repo *ordersRepository) CreateOrder(order entity.Orders) (entity.Orders, error) {
	err := order_repo.db.Create(&order).Error
	if err != nil {
		return entity.Orders{}, err
	}

	return order, nil
}

func (order_repo *ordersRepository) UpdateOrder(order entity.Orders) (entity.Orders, error) {
	err := order_repo.db.Save(&order).Error
	if err != nil {
		return entity.Orders{}, err
	}

	return order, nil
}

func (order_repo *ordersRepository) DeleteOrder(order entity.Orders) (entity.Orders, error) {
	err := order_repo.db.Delete(&order).Error
	if err != nil {
		return entity.Orders{}, nil
	}

	return order, nil
}
