package service

import (
	"errors"
	"time"

	"commercium/internal/entity"
	"commercium/internal/repository"
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
	ordersRepository   repository.OrdersRepository
	usersRepository    repository.UsersRepository
	productsRepository repository.ProductsRepository
}

func NewOrdersService(ordersRepo repository.OrdersRepository, usersRepo repository.UsersRepository, productsRepo repository.ProductsRepository) *ordersService {
	return &ordersService{
		ordersRepository:   ordersRepo,
		usersRepository:    usersRepo,
		productsRepository: productsRepo,
	}
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
		return nil, errors.New("date cannot be blank")
	}

	// VALIDASI UNTUK CEK APAKAH TO TIDAK LEBIH DULU DARI FROM
	if from.After(to) {
		return nil, errors.New("'from' date cannot be after 'to' date")
	}

	return order_serv.ordersRepository.GetOrderByDate(from, to)
}

func (order_serv *ordersService) CreateOrder(order entity.Orders) (entity.Orders, error) {

	// VALIDASI APAKAH PRODUCT ID VALID
	product, err := order_serv.productsRepository.GetProductByID(order.ProductID)
	if err != nil {
		return entity.Orders{}, errors.New("product not found")
	}

	// VALIDASI APAKAH USER ID VALID
	_, err = order_serv.usersRepository.GetUserByID(order.UserID)
	if err != nil {
		return entity.Orders{}, errors.New("user not found")
	}

	// VALIDASI UNTUK MENGECEK QUANTITY TIDAK 0
	if order.Quantity <= 0 {
		return entity.Orders{}, errors.New("quantity cannot be blank")
	}

	order.TotalPrice = float64(order.Quantity) * product.Price

	return order_serv.ordersRepository.CreateOrder(order)
}

func (order_serv *ordersService) UpdateOrder(id int, orderNew entity.Orders) (entity.Orders, error) {
	order, err := order_serv.ordersRepository.GetOrderByID(id)
	if err != nil {
		return entity.Orders{}, err
	}

	// VALIDASI APAKAH ATTRIBUT ORDER SUDAH DI INPUT
	if orderNew.Quantity != 0 {
		// VALIDASI UNTUK MENGECEK QUANTITY TIDAK 0
		if order.Quantity <= 0 {
			return entity.Orders{}, errors.New("quantity cannot be blank")
		}

		order.Quantity = orderNew.Quantity

		// MENGAMBIL HARGA PRODUCT
		product, err := order_serv.productsRepository.GetProductByID(order.ProductID)
		if err != nil {
			return entity.Orders{}, errors.New("product not found")
		}

		order.TotalPrice = float64(order.Quantity) * product.Price
	}
	if orderNew.Status != "" {
		order.Status = orderNew.Status
	}

	return order_serv.ordersRepository.UpdateOrder(order)
}

func (order_serv *ordersService) DeleteOrder(id int) (entity.Orders, error) {
	product, err := order_serv.ordersRepository.GetOrderByID(id)
	if err != nil {
		return entity.Orders{}, errors.New("order not found")
	}

	return order_serv.ordersRepository.DeleteOrder(product)
}
