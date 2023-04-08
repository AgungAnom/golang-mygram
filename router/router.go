package router

import (
	"golang-mygram/controllers"
	"golang-mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	//  Authentication
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	// Coment Router
	commentRouter := r.Group("/products")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/",controllers.CreateComment)
		commentRouter.PUT("/:commentID",middlewares.CommentAuthorization(),controllers.UpdateComment)
		commentRouter.GET("/:commentID",middlewares.CommentAuthorization(),controllers.GetComment)
		commentRouter.DELETE("/:commentID",middlewares.CommentAuthorization(),controllers.DeleteComment)

	}
	return r
} 