package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
ID 			uint	 	`gorm:"primaryKey" json:"id"`
UserID		uint		`gorm:"not null" json:"username" form:"username" validate:"required~full name required"`
PhotoID		string		`gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required~Email required,email~Invalid Email"`
Message		string		`gorm:"not null" json:"password" form:"password" validate:"required~Password required,MinStringLength(6)~Password has to have a minimum length of 6 characters"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`

}


func (c *Comment) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil{
		err = errCreate
		return	
	}

	return
}

