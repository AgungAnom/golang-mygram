package controllers

import (
	"fmt"
	"golang-mygram/database"
	"golang-mygram/helpers"
	"golang-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	SocialMedia := models.Socialmedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, SocialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	db := database.GetDB()
	OldSocialMedia := models.Socialmedia{}
	SocialMedia := models.Socialmedia{}
	SocialMediaID, _ := strconv.Atoi(c.Param("SocialMediaID"))
	userID := uint(userData["id"].(float64))



	err1 := db.First(&OldSocialMedia, SocialMediaID).Error
	if err1 != nil {
		c.AbortWithError(http.StatusInternalServerError, err1)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Social Media with id %v not found", SocialMediaID),
			})
		return
	}

	if err := c.ShouldBindJSON(&SocialMedia); err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	SocialMedia.UserID = userID
	SocialMedia.ID = uint(SocialMediaID)
	SocialMedia = models.Socialmedia{
		Name: SocialMedia.Name,
		SocialMediaURL: SocialMedia.SocialMediaURL,
	}

	err := db.Model(&SocialMedia).Where("id = ?", SocialMediaID).Updates(SocialMedia).Error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	err2 := db.First(&SocialMedia, SocialMediaID).Error
	if err2 != nil {
		c.AbortWithError(http.StatusInternalServerError, err2)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Social Media with id %v not found", SocialMediaID),
			})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)
}

func GetSocialMedia(c *gin.Context){
	SocialMediaID, _ := strconv.Atoi(c.Param("SocialMediaID"))
	SocialMedia := models.Socialmedia{}
	db := database.GetDB()

	err := db.First(&SocialMedia, SocialMediaID).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Social Media with id %v not found", SocialMediaID),
			})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context){
	SocialMediaID, _ := strconv.Atoi(c.Param("SocialMediaID"))
	SocialMedia := models.Socialmedia{}
	db := database.GetDB()

	err := db.First(&SocialMedia, SocialMediaID).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Social Media with id %v not found", SocialMediaID),
			})
		return
	}

	if err :=db.Delete(&SocialMedia).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Social Media deleted successfully",
	})
}