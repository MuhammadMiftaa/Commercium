package main

import (
	"commercium/config"
	"commercium/internal/entity"
	"commercium/internal/router"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}
	config.Migrate(db, &entity.Users{}, &entity.Products{}, &entity.Orders{})

	r := router.SetupRouter()
	r.Run(":8080")
}
