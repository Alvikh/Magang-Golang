package provinsi_v010

import "time"

type Provinsi struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Kode      string    `json:"kode" binding:"required" gorm:"type:char(2);unique"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(100);unique;not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
