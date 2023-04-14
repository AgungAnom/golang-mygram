package controllers

import (
	"errors"
	"golang-mygram/database"
	"golang-mygram/helpers"
	"golang-mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// UserRegister godoc
// @Summary Post new user
// @Description Post detail of a user
// @Tags users
// @Accept json
// @Produce json
// @Param models.User body models.User true "Register user"
// @Success 201 {object} models.User
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User :=models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id" : User.ID,
		"email": User.Email,
		"username":User.Username,
		"age" :User.Age,
	})
}

// UserLogin godoc
// @Summary Post to login account
// @Description Post email and password to login
// @Tags users
// @Accept json
// @Produce json
// @Param jsom body models.User true "Login user"
// @Success 200 {object} models.User
// @Router /users/login [post]
func UserLogin(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password


	err := db.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid password"))
		return
	}

	token, err := helpers.GenerateToken(User.ID, User.Email)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}