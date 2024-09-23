package group_v010

import "time"

type Group struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(100);unique;not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
