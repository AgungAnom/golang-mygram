package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represent the model for a comment
type Comment struct {
ID 			uint	 `gorm:"primaryKey" json:"id"`
UserID		uint		
PhotoID		uint	`gorm:"not null" json:"photo_id" form:"photo_id" valid:"required~Photo ID required"`
Message		string	`gorm:"not null" json:"message" form:"message" valid:"required~Message required"`
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

