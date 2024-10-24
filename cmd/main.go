package main

import (

	"commercium/config"
	"commercium/internal/entity"
	"commercium/interface/http/router"
	// "commercium/internal/repository"
	// "commercium/internal/service"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}
	config.Migrate(db, &entity.Users{}, &entity.Products{}, &entity.Orders{})

	// product_repo := repository.NewProductsRepository(db)
	// user_repo := repository.NewUsersRepository(db)
	// product_serv := service.NewProductsService(product_repo)
	// order_repo := repository.NewOrdersRepository(db)
	// order_serv := service.NewOrdersService(order_repo, user_repo, product_repo)

	// order := entity.Orders{
	// 	UserID: 5,
	// 	ProductID: 1,
	// 	Quantity: 24,
	// 	Status: "Paid",
	// }

	// order_serv.UpdateOrder(4, order)

	r := router.SetupRouter()
	r.Run(":8080")
}
