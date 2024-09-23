package user_v010

import "time"

type User struct {
	NIK       string    `json:"nik" binding:"required" gorm:"type:varchar(30);not null;primary_key"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
