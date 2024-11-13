package kecamatan_v010

import (
	"log"
	"rsudlampung/helper"
	"rsudlampung/modules/kabkot_v010"
	"rsudlampung/modules/provinsi_v010"

	"time"

	"gorm.io/gorm"
)

type KecamatanService interface {
	Create(Kecamatan) (Kecamatan, error)
	Update(Kecamatan) error
	Delete(Kecamatan) error
	FindAll() []Kecamatan
	FindById(kecamatanId uint64) Kecamatan
	FindByKabkot(kabkotId uint64) []Kecamatan
}

type kecamatanService struct {
	conn *gorm.DB
}

func NewKecamatanService(db *gorm.DB) KecamatanService {
	configEnv, errEnv := helper.LoadConfig("../.")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}
	am := configEnv.AutoMigrate

	// Auto migrate Kecamatan, Kabkot, and Provinsi
	if am == "on" {
		db.AutoMigrate(&Kecamatan{}, &kabkot_v010.Kabkot{}, &provinsi_v010.Provinsi{})
	}

	return &kecamatanService{
		conn: db,
	}
}

func (service *kecamatanService) Create(kecamatan Kecamatan) (Kecamatan, error) {
	kecamatan.CreatedAt = time.Now()
	err := service.conn.Create(&kecamatan).Error
	if err != nil {
		return Kecamatan{}, err
	}
	return kecamatan, nil
}

func (service *kecamatanService) Update(kecamatan Kecamatan) error {
	kecamatan.UpdatedAt = time.Now()
	err := service.conn.Save(&kecamatan).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *kecamatanService) Delete(kecamatan Kecamatan) error {
	err := service.conn.Delete(&kecamatan).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *kecamatanService) FindAll() []Kecamatan {
	var kecamatans []Kecamatan
	// Preload Kabkot and Provinsi
	service.conn.Preload("Kabkot").Preload("Kabkot.Provinsi").Find(&kecamatans)
	return kecamatans
}

func (service *kecamatanService) FindById(kecamatanId uint64) Kecamatan {
	var kecamatan Kecamatan
	// Preload Kabkot and Provinsi
	service.conn.Preload("Kabkot").Preload("Kabkot.Provinsi").Where("id = ?", kecamatanId).Find(&kecamatan)
	return kecamatan
}

func (service *kecamatanService) FindByKabkot(kabkotId uint64) []Kecamatan {
	var kecamatans []Kecamatan
	// Preload Kabkot and Provinsi
	service.conn.Preload("Kabkot").Preload("Kabkot.Provinsi").Where("kabkot_id = ?", kabkotId).Find(&kecamatans)
	return kecamatans
}
