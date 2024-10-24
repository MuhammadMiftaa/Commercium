package main

import (
	"commercium/interface/http/router"
	// "commercium/config"
	// "commercium/internal/entity"
)

func main() {
	// db, err := config.SetupDatabase()
	// if err != nil {
	// 	panic(err)
	// }
	// config.Migrate(db, &entity.Users{}, &entity.Products{}, &entity.Orders{})

	r := router.SetupRouter()
	r.Run(":8080")
}
