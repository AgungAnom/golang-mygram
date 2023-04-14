package models

import (
	"errors"
	"golang-mygram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
ID 			uint	 		`gorm:"primaryKey" json:"id"`
Username	string			`gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~username required"`
Email		string			`gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email required"`
Password	string			`gorm:"not null" json:"password" form:"password" valid:"required~Password required"`
Age			int				`gorm:"not null" json:"age" form:"age" valid:"required~Age must be numeric"`
Photo 		[]Photo			`gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
Comment		[]Comment		`gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
SocialMedia []Socialmedia	`gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
CreatedAt 	time.Time 		`json:"created_at"`
UpdatedAt 	time.Time 		`json:"updated_at"`

}


func (u *User) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil{
		err = errCreate
		return	
	}

	errEmail := govalidator.IsEmail(u.Email)
	if errEmail != true{
		err = errors.New("Invalid Email Address")
		return	
	}

	errPassword := len(u.Password)
	if errPassword <= 5{
		err = errors.New("Password has to have a minimum length of 6 characters")
		return	
	}

	errAge := u.Age
	if errAge < 8{
		err = errors.New("age must be above 8 years old")
		return	
	}


	hashedPass, err := helpers.HashPass(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass
	
	return
}



