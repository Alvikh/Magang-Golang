package kabkot_v010

import (
	"rsudlampung/modules/provinsi_v010"
	"time"
)

type Kabkot struct {
	ID         uint64                 `json:"id" gorm:"primary_key;auto_increment"`
	Name       string                 `json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	Provinsi   provinsi_v010.Provinsi `json:"provinsi" binding:"required" gorm:"foreignkey:ProvinsiID"`
	ProvinsiID uint64                 `json:"-"`
	CreatedAt  time.Time              `json:"-"`
	UpdatedAt  time.Time              `json:"-"`
	Status     bool                   `json:"status" gorm:"type:bool;default:false"`
	Fasprof    bool                   `json:"fasprof" gorm:"type:bool;default:false"`
	IN         bool                   `json:"in" gorm:"type:bool;default:false"`
}

type KabkotCreate struct {
	Name       string `json:"name" binding:"required"`
	ProvinsiID uint64 `json:"provinsi_id" binding:"required"`
	Status     bool   `json:"status" gorm:"type:bool;default:false"`
}

type KabkotEdit struct {
	ID         uint64 `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	ProvinsiID uint64 `json:"provinsi_id" binding:"required"`
	Fasprof    bool   `json:"fasprof"`
	IN         bool   `json:"in"`
	Status     bool   `json:"status" gorm:"type:bool;default:false"`
}
