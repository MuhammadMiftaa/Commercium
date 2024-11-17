package handler

import (
	"net/http"
	"strconv"

	"commercium/internal/entity"
	"commercium/internal/helper"
	"commercium/internal/service"

	"github.com/gin-gonic/gin"
)

type ordersHandler struct {
	ordersService service.OrdersService
}

func NewOrdersHandler(ordersService service.OrdersService) *ordersHandler {
	return &ordersHandler{ordersService}
}

func (order_handler *ordersHandler) GetAllOrders(c *gin.Context) {
	orders, err := order_handler.ordersService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	var ordersResponse []entity.OrdersResponse
	for _, order := range orders {
		orderResponse, _ := helper.ConvertToResponseType(order).(entity.OrdersResponse)
		ordersResponse = append(ordersResponse, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Get all orders data",
		"data":       orders,
	})
}

func (order_handler *ordersHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	idINT, _ := strconv.Atoi(id)

	order, err := order_handler.ordersService.GetOrderByID(idINT)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	orderResponse := helper.ConvertToResponseType(order)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Get order data",
		"data":       orderResponse,
	})
}

func (order_handler *ordersHandler) CreateOrder(c *gin.Context) {
	var orderRequest entity.OrdersRequest
	err := c.ShouldBindBodyWithJSON(&orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	order, err := order_handler.ordersService.CreateOrder(orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	orderResponse := helper.ConvertToResponseType(order)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Create order data",
		"data":       orderResponse,
	})
}

func (order_handler *ordersHandler) UpdateOrder(c *gin.Context) {
	var orderRequest entity.OrdersRequest
	err := c.ShouldBindBodyWithJSON(&orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	id := c.Param("id")
	idINT, _ := strconv.Atoi(id)

	product, err := order_handler.ordersService.UpdateOrder(idINT, orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	orderResponse := helper.ConvertToResponseType(product)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Update product data",
		"data":       orderResponse,
	})
}

func (order_handler *ordersHandler) PaidOrder(c *gin.Context) {
	id := c.Param("id")
	idINT, _ := strconv.Atoi(id)

	product, err := order_handler.ordersService.PaidOrder(idINT)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	orderResponse := helper.ConvertToResponseType(product)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Paid product data",
		"data":       orderResponse,
	})
}

func (order_handler *ordersHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	idINT, _ := strconv.Atoi(id)

	product, err := order_handler.ordersService.DeleteOrder(idINT)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	orderResponse := helper.ConvertToResponseType(product)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Delete product data",
		"data":       orderResponse,
	})
}
