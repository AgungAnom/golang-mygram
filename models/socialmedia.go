package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Socialmedia represent the model for a socialmedia
type Socialmedia struct {
ID 				uint	 	`gorm:"primaryKey" json:"id"`
Name			string		`gorm:"not null" json:"name" form:"name" valid:"required~Name required"`
SocialMediaURL	string		`gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social Media URL required"`
UserID			uint		
CreatedAt 		time.Time 	`json:"created_at"`
UpdatedAt 		time.Time 	`json:"updated_at"`
}


func (s *Socialmedia) BeforeCreate(projectDB *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil{
		err = errCreate
		return	
	}
	
	return
}

