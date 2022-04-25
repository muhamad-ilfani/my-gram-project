package routers

import (
	"my-gram/controllers"
	"my-gram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp(c controllers.UserRepo, p controllers.PhotoRepo) *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", c.GetAllUser)
		userRouter.POST("/register", c.UserRegister)
		userRouter.POST("/login", c.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), c.UserUpdate)
		userRouter.DELETE("/", middlewares.Authentication(), c.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", p.GetPhoto)
		photoRouter.POST("/", p.UploadPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), p.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), p.DeletePhoto)
	}
	return r
}
