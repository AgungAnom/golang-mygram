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
		socialMediaRouter.POST("/",controllers.CreateSocialMedia)
		socialMediaRouter.GET("/",controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/:socialMediaID",controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaID",middlewares.SocialMediaAuthorization(),controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaID",middlewares.SocialMediaAuthorization(),controllers.DeleteSocialMedia)
	}

	// Photo Router
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/",controllers.CreatePhoto)
		photoRouter.GET("/",controllers.GetAllPhoto)
		photoRouter.GET("/:photoID",controllers.GetPhoto)
		photoRouter.PUT("/:photoID",middlewares.PhotoAuthorization(),controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID",middlewares.PhotoAuthorization(),controllers.DeletePhoto)
	}
	// Comment Router
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/",controllers.CreateComment)
		commentRouter.GET("/:commentID",controllers.GetComment)
		commentRouter.PUT("/:commentID",middlewares.CommentAuthorization(),controllers.UpdateComment)
		commentRouter.DELETE("/:commentID",middlewares.CommentAuthorization(),controllers.DeleteComment)
	}
	return r
} 