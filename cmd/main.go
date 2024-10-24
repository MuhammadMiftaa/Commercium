package main

import (
	"commercium/config"
	"commercium/interface/http/router"
	// "commercium/internal/entity"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}
	// config.Migrate(db, &entity.Users{}, &entity.Products{}, &entity.Orders{})

	r := router.SetupRouter(db)
	r.Run(":8080")
}
