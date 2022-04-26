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
	commentRepo := controllers.CommentRepo{DB: db}
	mediaRepo := controllers.MediaRepo{DB: db}

	r := routers.StartApp(userRepo, photoRepo, commentRepo, mediaRepo)
	r.Run(":8080")
}
