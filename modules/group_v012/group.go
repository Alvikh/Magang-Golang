package group_v012

import "time"

type Group struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(100);unique;not null"`
	Scope     string    `json:"scope" binding:"required" gorm:"type:varchar(100);unique;not null"`
	Domain    string    `json:"domain" binding:"required" gorm:"type:varchar(100);unique;not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
