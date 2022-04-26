package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Media struct {
	GormModel
	Name             string `gorm:"not null" json:"name_m" form:"name_m" valid:"required~Name is required"`
	Social_media_url string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social_media_url is required"`
	User_id          uint   `gorm:"user_id"`
	User             *User
}

func (m *Media) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (m *Media) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
