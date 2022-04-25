package main

import (
	"my-gram/config"
	"my-gram/controllers"
	"my-gram/routers"
)

func main() {
	db := config.InitDB()

	userRepo := controllers.UserRepo{DB: db}
	photoRepo := controllers.PhotoRepo{DB: db}

	r := routers.StartApp(userRepo, photoRepo)
	r.Run(":8080")
}
