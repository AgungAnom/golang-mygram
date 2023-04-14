package middlewares

import (
	"golang-mygram/database"
	"golang-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoID, err := strconv.Atoi(c.Param("photoID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Parameter",
			})
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo,uint(photoID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data Not Found",
				"message" : "Data doesn't exist",
			})
			return
		}

		if Photo.UserID != userID{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Unauthorized",
				"message" : "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentID, err := strconv.Atoi(c.Param("commentID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Parameter",
			})
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment,uint(commentID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data Not Found",
				"message" : "Data doesn't exist",
			})
			return
		}

		if Comment.UserID != userID{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Unauthorized",
				"message" : "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaID, err := strconv.Atoi(c.Param("socialMediaID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Parameter",
			})
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.Socialmedia{}

		err = db.Select("user_id").First(&SocialMedia,uint(socialMediaID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data Not Found",
				"message" : "Data doesn't exist",
			})
			return
		}

		if SocialMedia.UserID != userID{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Unauthorized",
				"message" : "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}