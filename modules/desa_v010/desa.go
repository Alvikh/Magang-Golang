package desa_v010

import (
	"rsudlampung/modules/kecamatan_v010"
	"time"
)

type Desa struct {
	ID          uint64                   `json:"id" gorm:"primary_key;auto_increment"`
	Name        string                   `json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	Kecamatan   kecamatan_v010.Kecamatan `json:"kecamatan" binding:"required" gorm:"foreignkey:KecamatanID"`
	KecamatanID uint64                   `json:"kecamatan_id" binding:"required"`
	CreatedAt   time.Time                `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time                `json:"updated_at" gorm:"autoUpdateTime"`
}

type DesaCreate struct {
	Name        string `json:"name" binding:"required"`
	KecamatanID uint64 `json:"kecamatan_id" binding:"required"`
}

type DesaEdit struct {
	ID          uint64 `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	KecamatanID uint64 `json:"kecamatan_id" binding:"required"`
}
