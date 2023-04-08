package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
ID 				uint	 	`gorm:"primaryKey" json:"id"`
Name			string		`gorm:"not null" json:"username" form:"username" validate:"required~full name required"`
SocialMediaURL	string		`gorm:"not null" json:"password" form:"password" validate:"required~Password required,MinStringLength(6)~Password has to have a minimum length of 6 characters"`
UserID			uint		`gorm:"not null" json:"age" form:"age" validate:"required~Age required"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`

}


func (s *SocialMedia) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil{
		err = errCreate
		return	
	}
	
	return
}

