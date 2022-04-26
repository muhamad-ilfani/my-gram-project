package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	User_id  uint   `gorm:"user_id"`
	Photo_id uint   `gorm:"photo_id" json:"photo_id" form:"photo_id"`
	Message  string `gorm:"not null" json:"message" form:"message" valid:"required~message is required"`
	User     *User
	Photo    *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
