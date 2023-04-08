package models

import (
	"golang-mygram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
ID 			uint	 	`gorm:"primaryKey" json:"id"`
UserName	string		`gorm:"not null" json:"username" form:"username" validate:"required~full name required"`
Email		string		`gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required~Email required,email~Invalid Email"`
Password	string		`gorm:"not null" json:"password" form:"password" validate:"required~Password required,MinStringLength(6)~Password has to have a minimum length of 6 characters"`
Age			int			`gorm:"not null" json:"age" form:"age" validate:"required~Age required"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`

}


func (u *User) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil{
		err = errCreate
		return	
	}
	hashedPass, err := helpers.HashPass(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass
	
	return
}

