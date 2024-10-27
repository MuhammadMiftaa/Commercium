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

	r := router.SetupRouter(db)
	r.Run(":8080")
}
