package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
ID 				uint	 	`gorm:"primaryKey" json:"id"`
Name			string		`gorm:"not null" json:"name" form:"name" validate:"required~Name required"`
SocialMediaURL	string		`gorm:"not null" json:"social_media_url" form:"social_media_url" validate:"required~Social Media URL required"`
UserID			uint		
CreatedAt 		time.Time 	`json:"created_at"`
UpdatedAt 		time.Time 	`json:"updated_at"`
}


func (s *SocialMedia) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil{
		err = errCreate
		return	
	}
	
	return
}

