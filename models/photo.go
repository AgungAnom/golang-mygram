package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
ID 			uint	 	`gorm:"primaryKey" json:"id"`
Title		string		`gorm:"not null" json:"username" form:"username" validate:"required~full name required"`
Caption		string		`gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required~Email required,email~Invalid Email"`
PhotoURL	string		`gorm:"not null" json:"password" form:"password" validate:"required~Password required,MinStringLength(6)~Password has to have a minimum length of 6 characters"`
UserID		uint		`gorm:"not null" json:"age" form:"age" validate:"required~Age required"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`

}


func (p *Photo) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil{
		err = errCreate
		return	
	}
	return
}

