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

// CreatePhoto godoc
// @Summary Post new photo
// @Description Post detail of a photo
// @Tags photos
// @Accept json
// @Produce json
// @Param models.Photo body models.Photo true "Create photo"
// @Success 201 {object} models.Photo
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
}

// UpdatePhoto godoc
// @Summary Put update to photo identified by id
// @Description Put update detail of a photo corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param photoID path uint true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photos/{photoID} [put]
func UpdatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	db := database.GetDB()
	OldPhoto := models.Photo{}
	Photo := models.Photo{}
	PhotoID, _ := strconv.Atoi(c.Param("photoID"))
	userID := uint(userData["id"].(float64))



	err1 := db.First(&OldPhoto, PhotoID).Error
	if err1 != nil {
		c.AbortWithError(http.StatusInternalServerError, err1)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Photo with id %v not found", PhotoID),
			})
		return
	}

	if err := c.ShouldBindJSON(&Photo); err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	Photo.UserID = userID
	Photo.ID = uint(PhotoID)
	Photo = models.Photo{
		Title: Photo.Title,
		Caption: Photo.Caption,
		PhotoURL: Photo.PhotoURL,
	}

	err := db.Model(&Photo).Where("id = ?", PhotoID).Updates(Photo).Error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	err2 := db.First(&Photo, PhotoID).Error
	if err2 != nil {
		c.AbortWithError(http.StatusInternalServerError, err2)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Photo with id %v not found", PhotoID),
			})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

// GetPhoto godoc
// @Summary Get a photo detail by id
// @Description Get data of photo corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param photoID path uint true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photos/{photoID} [get]
func GetPhoto(c *gin.Context){
	PhotoID, _ := strconv.Atoi(c.Param("photoID"))
	Photo := models.Photo{}
	db := database.GetDB()

	err := db.First(&Photo, PhotoID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error" :"Data Not Found",
			"message": fmt.Sprintf("Photo with id %v not found", PhotoID),
			})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto godoc
// @Summary Delete a photo detail by id
// @Description Delete data of photo corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param photoID path uint true "ID of the photo"
// @Success 200 "Photo successfully deleted"
// @Router /photos/{photoID} [delete]
func DeletePhoto(c *gin.Context){
	PhotoID, _ := strconv.Atoi(c.Param("photoID"))
	Photo := models.Photo{}
	db := database.GetDB()

	err := db.First(&Photo, PhotoID).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Photo with id %v not found", PhotoID),
			})
		return
	}

	if err :=db.Delete(&Photo).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Photo successfully deleted",
	})
}

// GetAllPhoto godoc
// @Summary Get details
// @Description Get data of all photo
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos [get]
func GetAllPhoto(c *gin.Context){
	db := database.GetDB()
	Photo := []models.Photo{}
	err := db.Find(&Photo).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK,Photo)
}