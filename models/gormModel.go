package models

import "time"

type GormModel struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}
