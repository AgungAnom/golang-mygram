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

func CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	// Check photo before create comment
	Photo := models.Photo{}
	errPhoto := db.First(&Photo,  Comment.PhotoID).Error
	if errPhoto != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": "Photo not found",
			})
		return
	}
	
	Comment.UserID = userID
	err := db.Debug().Create(&Comment).Error
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Comment)
}

func UpdateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	db := database.GetDB()
	OldComment := models.Comment{}
	Comment := models.Comment{}
	CommentID, _ := strconv.Atoi(c.Param("commentID"))
	userID := uint(userData["id"].(float64))



	err1 := db.First(&OldComment, CommentID).Error
	if err1 != nil {
		c.AbortWithError(http.StatusInternalServerError, err1)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Comment with id %v not found", CommentID),
			})
		return
	}

	if err := c.ShouldBindJSON(&Comment); err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Check photo before update comment
	Photo := models.Photo{}
	errPhoto := db.First(&Photo,  Comment.PhotoID).Error
	if errPhoto != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": "Photo not found",
			})
		return
	}


	Comment.UserID = userID
	Comment.ID = uint(CommentID)
	Comment = models.Comment{
		PhotoID: Comment.PhotoID,
		Message: Comment.Message,
	}

	err := db.Model(&Comment).Where("id = ?", CommentID).Updates(Comment).Error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	err2 := db.First(&Comment, CommentID).Error
	if err2 != nil {
		c.AbortWithError(http.StatusInternalServerError, err2)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Comment with id %v not found", CommentID),
			})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

func GetComment(c *gin.Context){
	CommentID, _ := strconv.Atoi(c.Param("commentID"))
	Comment := models.Comment{}
	db := database.GetDB()

	err := db.First(&Comment, CommentID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error" :"Data Not Found",
			"message": fmt.Sprintf("Comment with id %v not found", CommentID),
			})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context){
	CommentID, _ := strconv.Atoi(c.Param("commentID"))
	Comment := models.Comment{}
	db := database.GetDB()

	err := db.First(&Comment, CommentID).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Comment with id %v not found", CommentID),
			})
		return
	}

	if err :=db.Delete(&Comment).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Comment deleted successfully",
	})
}

func GetAllComment(c *gin.Context){
	db := database.GetDB()
	Comment := []models.Comment{}
	err := db.Find(&Comment).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK,Comment)
}