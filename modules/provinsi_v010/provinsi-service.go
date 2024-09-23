package provinsi_v010

import (
	"log"
	"rsudlampung/helper"
	"time"

	"gorm.io/gorm"
)

type ProvinsiService interface {
	Create(Provinsi) (Provinsi, error)
	Update(Provinsi) error
	Delete(Provinsi) error
	FindAll() []Provinsi
	FindById(provinsiId uint64) Provinsi
}

type provinsiService struct {
	conn *gorm.DB
}

func NewProvinsiService(db *gorm.DB) ProvinsiService {
	configEnv, errEnv := helper.LoadConfig("../.")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}
	am := configEnv.AutoMigrate

	if am == "on" {
		db.AutoMigrate(&Provinsi{})
	}

	return &provinsiService{
		conn: db,
	}
}

func (service *provinsiService) Create(provinsi Provinsi) (Provinsi, error) {
	provinsi.UpdatedAt = time.Now()
	err := service.conn.Create(&provinsi).Error
	if err != nil {
		return Provinsi{}, err
	}
	return provinsi, nil

}

func (service *provinsiService) Update(provinsi Provinsi) error {
	provinsi.UpdatedAt = time.Now()
	err := service.conn.Save(&provinsi).Error
	if err != nil {
		return err
	}
	return nil

}

func (service *provinsiService) Delete(provinsi Provinsi) error {
	err := service.conn.Delete(&provinsi).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *provinsiService) FindAll() []Provinsi {
	var provinsis []Provinsi
	service.conn.Find(&provinsis)
	return provinsis

}

func (service *provinsiService) FindById(provinsiId uint64) Provinsi {
	var provinsi Provinsi
	service.conn.Where("id=?", provinsiId).Find(&provinsi)
	return provinsi

}
