package handler

import (
	"net/http"
	"strconv"

	"commercium/internal/entity"
	"commercium/internal/helper"
	"commercium/internal/service"

	"github.com/gin-gonic/gin"
)

type productsHandler struct {
	productsService service.ProductsService
}

func NewProductsHandler(productsService service.ProductsService) *productsHandler {
	return &productsHandler{productsService}
}

func (product_handler *productsHandler) GetAllProducts(c *gin.Context) {
	products, err := product_handler.productsService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	var productsResponse []entity.ProductsResponse
	for _, product := range products {
		productResponse, _ := helper.ConvertToResponseType(product).(entity.ProductsResponse)
		productsResponse = append(productsResponse, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Get all products data",
		"data":       productsResponse,
	})
}

func (product_handler *productsHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	idINT, _ := strconv.Atoi(id)

	product, err := product_handler.productsService.GetProductByID(idINT)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	productResponse := helper.ConvertToResponseType(product)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Get product data",
		"data":       productResponse,
	})
}

func (product_handler *productsHandler) CreateProduct(c *gin.Context) {
	var productRequest entity.ProductsRequest
	err := c.ShouldBindBodyWithJSON(&productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	product, err := product_handler.productsService.CreateProduct(productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	productResponse := helper.ConvertToResponseType(product)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Create product data",
		"data":       productResponse,
	})
}

func (product_handler *productsHandler) UpdateProduct(c *gin.Context) {
	var productRequest entity.ProductsRequest
	err := c.ShouldBindBodyWithJSON(&productRequest)
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

	product, err := product_handler.productsService.UpdateProduct(idINT, productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	productResponse := helper.ConvertToResponseType(product)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Update product data",
		"data":       productResponse,
	})
}

func (product_handler *productsHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	idINT, _ := strconv.Atoi(id)

	product, err := product_handler.productsService.DeleteProduct(idINT)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
		return
	}

	// MENGUBAH TIPE ENITITY KE TIPE RESPONSE
	productResponse := helper.ConvertToResponseType(product)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"status":     true,
		"message":    "Delete product data",
		"data":       productResponse,
	})
}
