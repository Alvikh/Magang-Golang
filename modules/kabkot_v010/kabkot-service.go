package kabkot_v010

import (
	"log"
	"rsudlampung/helper"
	"time"

	"gorm.io/gorm"
)

type KabkotService interface {
	Create(Kabkot) (Kabkot, error)
	Update(Kabkot) error
	Delete(Kabkot) error
	FindAll() []Kabkot
	FindById(kabkotId uint64) Kabkot
	FindByProvinsi(provinsiId uint64) []Kabkot
}

type kabkotService struct {
	conn *gorm.DB
}

func NewKabkotService(db *gorm.DB) KabkotService {
	configEnv, errEnv := helper.LoadConfig("../.")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}
	am := configEnv.AutoMigrate

	if am == "on" {
		db.AutoMigrate(&Kabkot{})
	}

	return &kabkotService{
		conn: db,
	}
}

func (service *kabkotService) Create(kabkot Kabkot) (Kabkot, error) {
	kabkot.CreatedAt = time.Now()
	err := service.conn.Create(&kabkot).Error
	if err != nil {
		return Kabkot{}, err
	}
	return kabkot, nil
}

func (service *kabkotService) Update(kabkot Kabkot) error {
	kabkot.UpdatedAt = time.Now()
	err := service.conn.Save(&kabkot).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *kabkotService) Delete(kabkot Kabkot) error {
	err := service.conn.Delete(&kabkot).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *kabkotService) FindAll() []Kabkot {
	var kabkots []Kabkot
	service.conn.Preload("Provinsi").Find(&kabkots)
	return kabkots
}

func (service *kabkotService) FindById(kabkotId uint64) Kabkot {
	var kabkot Kabkot
	service.conn.Preload("Provinsi").Where("id=?", kabkotId).Find(&kabkot)
	return kabkot
}

func (service *kabkotService) FindByProvinsi(provinsiId uint64) []Kabkot {
	var kabkots []Kabkot
	service.conn.Preload("Provinsi").Where("provinsi_id=?", provinsiId).Find(&kabkots)
	return kabkots
}
