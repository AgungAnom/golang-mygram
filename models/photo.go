package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
ID 			uint	 	`gorm:"primaryKey" json:"id"`
Title		string		`gorm:"not null" json:"title" form:"title" validate:"required~Title required"`
Caption		string		`json:"caption" form:"caption"`
PhotoURL	string		`gorm:"not null" json:"photo_url" form:"password" validate:"required~Photo URL required"`
UserID		uint		
Comment []Comment		`gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
CreatedAt time.Time 	`json:"created_at"`
UpdatedAt time.Time 	`json:"updated_at"`
}


func (p *Photo) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil{
		err = errCreate
		return	
	}
	return
}

