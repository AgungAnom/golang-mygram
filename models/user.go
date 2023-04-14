package models

import (
	"golang-mygram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
ID 			uint	 		`gorm:"primaryKey" json:"id"`
Username	string			`gorm:"not null;uniqueIndex" json:"username" form:"username" validate:"required~username required"`
Email		string			`gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required~Email required,email~Invalid email address"`
Password	string			`gorm:"not null" json:"password" form:"password" validate:"required~Password required,MinStringLength(6)~Password has to have a minimum length of 6 characters"`
Age			int				`gorm:"not null" json:"age" form:"age" validate:"required~Age required, numeric~Age must be numeric,min=9~age must be above 9 years old"`
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

	hashedPass, err := helpers.HashPass(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass
	
	return
}

