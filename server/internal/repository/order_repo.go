package repository

import (
	"errors"
	"time"

	"commercium/internal/entity"

	"gorm.io/gorm"
)

type OrdersRepository interface {
	GetAllOrders() ([]entity.OrderDetail, error)
	GetOrderByID(id int) (entity.Orders, error)
	GetOrderByDate(from, to time.Time) ([]entity.Orders, error)
	CreateOrder(order entity.Orders) (entity.Orders, error)
	UpdateOrder(order entity.Orders) (entity.Orders, error)
	PaidOrder(id int, user_id int, product_id int) (entity.Orders, error)
	DeleteOrder(order entity.Orders) (entity.Orders, error)
}

type ordersRepository struct {
	db *gorm.DB
}

func NewOrdersRepository(db *gorm.DB) *ordersRepository {
	return &ordersRepository{db}
}

func (order_repo *ordersRepository) GetAllOrders() ([]entity.OrderDetail, error) {
	var orders []entity.OrderDetail
	err := order_repo.db.Table("orders").
		Select("orders.id AS id, users.fullname AS customer_name, products.name AS product_name, orders.quantity AS quantity, products.price AS product_price, orders.quantity * products.price AS total_price, orders.status AS status").
		Joins("INNER JOIN users ON users.id = orders.user_id").
		Joins("INNER JOIN products ON products.id = orders.product_id").
		Order("id ASC").
		Scan(&orders).
		Error
	if err != nil {
		return nil, errors.New("failed to get orders")
	}

	return orders, nil
}

func (order_repo *ordersRepository) GetOrderByID(id int) (entity.Orders, error) {
	var order entity.Orders
	err := order_repo.db.First(&order, id).Error
	if err != nil {
		return entity.Orders{}, errors.New("order not found")
	}

	return order, nil
}

func (order_repo *ordersRepository) GetOrderByDate(from, to time.Time) ([]entity.Orders, error) {
	var orders []entity.Orders
	err := order_repo.db.Where("created_at BETWEEN ? AND ?", from, to).Find(&orders).Error
	if err != nil {
		return nil, errors.New("failed to get orders")
	}

	return orders, nil
}

func (order_repo *ordersRepository) CreateOrder(order entity.Orders) (entity.Orders, error) {
	err := order_repo.db.Create(&order).Error
	if err != nil {
		return entity.Orders{}, errors.New("failed to create order")
	}

	return order, nil
}

func (order_repo *ordersRepository) UpdateOrder(order entity.Orders) (entity.Orders, error) {
	err := order_repo.db.Save(&order).Error
	if err != nil {
		return entity.Orders{}, errors.New("failed to update order")
	}

	return order, nil
}

func (order_repo *ordersRepository) PaidOrder(id int, user_id int, product_id int) (entity.Orders, error) {
    // Mulai transaksi
    tx := order_repo.db.Begin()
    if tx.Error != nil {
        return entity.Orders{}, errors.New("failed to begin transaction")
    }

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback() // Rollback jika terjadi panic
        }
    }()

    var order entity.Orders
    var product entity.Products
    var user entity.Users
    var admin entity.Users

    // Logika bisnis
    if err := tx.First(&order, id).Error; err != nil {
        tx.Rollback()
        return entity.Orders{}, errors.New("order not found")
    }

    if err := tx.First(&product, product_id).Error; err != nil {
        tx.Rollback()
        return entity.Orders{}, errors.New("product not found")
    }

    if err := tx.First(&user, user_id).Error; err != nil {
        tx.Rollback()
        return entity.Orders{}, errors.New("user not found")
    }

    if err := tx.Where("role = ?", "admin").First(&admin).Error; err != nil {
        tx.Rollback()
        return entity.Orders{}, errors.New("admin not found")
    }

    if order.Status == "paid" {
        tx.Rollback()
        return entity.Orders{}, errors.New("order already paid")
    }

    if order.Quantity > product.Stock {
        tx.Rollback()
        return entity.Orders{}, errors.New("insufficient stock")
    }

    if user.Role == "admin" {
        tx.Rollback()
        return entity.Orders{}, errors.New("admin cannot make a payment")
    }

    // Update data
    product.Stock -= order.Quantity
    admin.Income += order.TotalPrice
    user.Outcome += order.TotalPrice
    order.Status = "paid"

    // Simpan data
    if err := tx.Save(&product).Error; err != nil {
        tx.Rollback()
        return entity.Orders{}, errors.New("failed to update product")
    }

    if err := tx.Save(&admin).Error; err != nil {
        tx.Rollback()
        return entity.Orders{}, errors.New("failed to update admin")
    }

    if err := tx.Save(&user).Error; err != nil {
        tx.Rollback()
        return entity.Orders{}, errors.New("failed to update user")
    }

    if err := tx.Save(&order).Error; err != nil {
        tx.Rollback()
        return entity.Orders{}, errors.New("failed to update order")
    }

    // Commit transaksi
    if err := tx.Commit().Error; err != nil {
        return entity.Orders{}, errors.New("failed to commit transaction")
    }

    return order, nil
}


func (order_repo *ordersRepository) DeleteOrder(order entity.Orders) (entity.Orders, error) {
	err := order_repo.db.Delete(&order).Error
	if err != nil {
		return entity.Orders{}, errors.New("failed to delete order")
	}

	return order, nil
}
