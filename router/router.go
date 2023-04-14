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

	// SocialMedia Router
	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/",controllers.CreateComment)
		socialMediaRouter.PUT("/:socialMediaID",middlewares.Authorization(),controllers.UpdateSocialMedia)
		socialMediaRouter.GET("/:socialMediaID",middlewares.Authorization(),controllers.GetSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaID",middlewares.Authorization(),controllers.DeleteSocialMedia)
	}

	// Photo Router
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/",controllers.CreateComment)
		photoRouter.PUT("/:photoID",middlewares.Authorization(),controllers.UpdatePhoto)
		photoRouter.GET("/:photoID",middlewares.Authorization(),controllers.GetPhoto)
		photoRouter.DELETE("/:photoID",middlewares.Authorization(),controllers.DeletePhoto)
	}
	// Comment Router
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/",controllers.CreateComment)
		commentRouter.PUT("/:commentID",middlewares.Authorization(),controllers.UpdateComment)
		commentRouter.GET("/:commentID",middlewares.Authorization(),controllers.GetComment)
		commentRouter.DELETE("/:commentID",middlewares.Authorization(),controllers.DeleteComment)
	}
	return r
} 