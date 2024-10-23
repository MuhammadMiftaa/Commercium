package service

import (
	"commercium/internal/entity"
	"commercium/internal/repository"
	"errors"
	"time"
)

type OrdersService interface {
	GetAllOrders() ([]entity.Orders, error)
	GetOrderByID(id int) (entity.Orders, error)
	GetOrderByDate(from, to time.Time) ([]entity.Orders, error)
	CreateOrder(order entity.Orders) (entity.Orders, error)
	UpdateOrder(id int, orderNew entity.Orders) (entity.Orders, error)
	DeleteOrder(id int) (entity.Orders, error)
}

type ordersService struct {
	ordersRepository repository.OrdersRepository
}

func NewOrdersService(ordersRepository repository.OrdersRepository) *ordersService {
	return &ordersService{ordersRepository}
}

func (order_serv *ordersService) GetAllOrders() ([]entity.Orders, error) {
	return order_serv.ordersRepository.GetAllOrders()
}

func (order_serv *ordersService) GetOrderByID(id int) (entity.Orders, error) {
	return order_serv.ordersRepository.GetOrderByID(id)
}

func (order_serv *ordersService) GetOrderByDate(from, to time.Time) ([]entity.Orders, error) {
	// VALIDASI UNTUK CEK APAKAH FROM DAN TO TIDAK KOSONG
	if from.IsZero() || to.IsZero() {
		return nil, errors.New("Date cannot be blank")
	}

	// VALIDASI UNTUK CEK APAKAH TO TIDAK LEBIH DULU DARI FROM
	if from.After(to) {
		return nil, errors.New("'from' date cannot be after 'to' date")
	}

	return order_serv.ordersRepository.GetOrderByDate(from, to)
}

func (order_serv *ordersService) CreateOrder(order entity.Orders) (entity.Orders, error) {
	order.TotalPrice = float64(order.Quantity) * 100000

	return order_serv.ordersRepository.CreateOrder(order)
}

func (order_serv *ordersService) UpdateOrder(id int, orderNew entity.Orders) (entity.Orders, error) {
	order, err := order_serv.ordersRepository.GetOrderByID(id)
	if err != nil {
		return orderNew, err
	}

	// VALIDASI APAKAH ATTRIBUT ORDER SUDAH DI INPUT
	if orderNew.Quantity != 0 {
		order.Quantity = orderNew.Quantity
		order.TotalPrice = float64(order.Quantity) * 100000
	}
	if orderNew.Status != "" {
		order.Status = orderNew.Status
	}

	return order_serv.ordersRepository.DeleteOrder(order)
}

func (order_serv *ordersService) DeleteOrder(id int) (entity.Orders, error) {
	product, err := order_serv.ordersRepository.GetOrderByID(id)
	if err != nil {
		return product, errors.New("Order not found")
	}

	return order_serv.ordersRepository.DeleteOrder(product)
}