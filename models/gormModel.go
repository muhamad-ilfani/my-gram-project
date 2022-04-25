package models

import "time"

type GormModel struct {
	Id         int       `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}
