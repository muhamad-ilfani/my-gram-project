package main

import (
	"my-gram/config"
	"my-gram/controllers"
	"my-gram/routers"
)

func main() {
	db := config.InitDB()

	InDB := controllers.UserRepo{DB: db}

	r := routers.StartApp(InDB)
	r.Run(":8080")
}
