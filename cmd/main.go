package main

import (
	"commercium/config"
	"commercium/internal/entity"
	"commercium/internal/repository"
	"commercium/internal/router"
	"commercium/internal/service"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}
	config.Migrate(db, &entity.Users{}, &entity.Products{}, &entity.Orders{})

	product_repo := repository.NewProductsRepository(db)
	product_serv := service.NewProductsService(product_repo)
	product_serv.DeleteProduct(2)

	r := router.SetupRouter()
	r.Run(":8080")
}
