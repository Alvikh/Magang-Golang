package kecamatan_v010

import (
	"rsudlampung/modules/kabkot_v010"
	"time"
)

type Kecamatan struct {
	ID        uint64             `json:"id" gorm:"primary_key;auto_increment"`
	Name      string             `json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	Kabkot    kabkot_v010.Kabkot `json:"kabkot" binding:"required" gorm:"foreignkey:KabkotID"`
	KabkotID  uint64             `json:"kabkot_id"`
	CreatedAt time.Time          `json:"-"`
	UpdatedAt time.Time          `json:"-"`
}

type KecamatanCreate struct {
	Name     string `json:"name" binding:"required"`
	KabkotID uint64 `json:"kabkot_id" binding:"required"`
}

type KecamatanEdit struct {
	ID       uint64 `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	KabkotID uint64 `json:"kabkot_id" binding:"required"`
}
