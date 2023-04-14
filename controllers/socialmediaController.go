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

// CreateSocialMedia godoc
// @Summary Post new social-media
// @Description Post detail of a social media
// @Tags social-media
// @Accept json
// @Produce json
// @Param models.Socialmedia body models.Socialmedia true "Create social media"
// @Success 201 {object} models.Socialmedia
// @Router /social-media [post]
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

// UpdateSocialMedia godoc
// @Summary Put update to social media identified by id
// @Description Put update detail of a social media corresponding to the input id
// @Tags social-media
// @Accept json
// @Produce json
// @Param socialMediaID path uint true "ID of the social media"
// @Success 200 {object} models.Socialmedia
// @Router /social-media/{socialMediaID} [put]
func UpdateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	db := database.GetDB()
	OldSocialMedia := models.Socialmedia{}
	SocialMedia := models.Socialmedia{}
	SocialMediaID, _ := strconv.Atoi(c.Param("socialMediaID"))
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
// GetSocialMedia godoc
// @Summary Get a social media detail by id
// @Description Get data of social media corresponding to the input id
// @Tags social-media
// @Accept json
// @Produce json
// @Param socialMediaID path uint true "ID of the social media"
// @Success 200 {object} models.Socialmedia
// @Router /social-media/{socialMediaID} [get]
func GetSocialMedia(c *gin.Context){
	SocialMediaID, _ := strconv.Atoi(c.Param("socialMediaID"))
	SocialMedia := models.Socialmedia{}
	db := database.GetDB()

	err := db.First(&SocialMedia, SocialMediaID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error" :"Data Not Found",
			"message": fmt.Sprintf("Social Media with id %v not found", SocialMediaID),
			})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)
}

// DeleteSocialMedia godoc
// @Summary Delete a social media detail by id
// @Description Delete data of social media corresponding to the input id
// @Tags social-media
// @Accept json
// @Produce json
// @Param socialMediaID path uint true "ID of the social media"
// @Success 200 "Social Media successfully deleted"
// @Router /social-media/{socialMediaID} [delete]
func DeleteSocialMedia(c *gin.Context){
	SocialMediaID, _ := strconv.Atoi(c.Param("socialMediaID"))
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
		"message":"Social Media successfully deleted",
	})
}

// GetAllSocialMedia godoc
// @Summary Get details
// @Description Get data of all social media
// @Tags social-media
// @Accept json
// @Produce json
// @Success 200 {object} models.Socialmedia
// @Router /social-media [get]
func GetAllSocialMedia(c *gin.Context){
	db := database.GetDB()
	SocialMedia := []models.Socialmedia{}
	err := db.Find(&SocialMedia).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK,SocialMedia)
}