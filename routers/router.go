package routers

import (
	"my-gram/controllers"
	"my-gram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp(c controllers.UserRepo, p controllers.PhotoRepo, o controllers.CommentRepo, m controllers.MediaRepo) *gin.Engine {
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

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", o.GetComment)
		commentRouter.POST("/", o.UploadComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), o.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), o.DeleteComment)
	}

	mediaRouter := r.Group("/socialmedias")
	{
		mediaRouter.Use(middlewares.Authentication())
		mediaRouter.GET("/", m.GetMedia)
		mediaRouter.POST("/", m.UploadMedia)
		mediaRouter.PUT("/:socialMediaId", middlewares.MediaAuthorization(), m.UpdateMedia)
		mediaRouter.DELETE("/:socialMediaId", middlewares.MediaAuthorization(), m.DeleteMedia)
	}

	return r
}
