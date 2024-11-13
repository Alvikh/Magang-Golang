package kecamatan_v010

import (
	"errors"
	"rsudlampung/modules/kabkot_v010"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type KecamatanController interface {
	FindAll(ctx *gin.Context) []Kecamatan
	FindByKabkot(ctx *gin.Context) []Kecamatan
	Create(ctx *gin.Context) (Kecamatan, error)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controllerKecamatan struct {
	service       KecamatanService
	serviceKabkot kabkot_v010.KabkotService
}

// NewKecamatanController initializes the KecamatanController
func NewKecamatanController(db *gorm.DB) KecamatanController {
	return &controllerKecamatan{
		service:       NewKecamatanService(db),
		serviceKabkot: kabkot_v010.NewKabkotService(db),
	}
}

func (c *controllerKecamatan) FindAll(ctx *gin.Context) []Kecamatan {
	return c.service.FindAll()
}

func (c *controllerKecamatan) FindByKabkot(ctx *gin.Context) []Kecamatan {
	kabkotId, err := strconv.ParseUint(ctx.Param("kabkot_id"), 10, 64)
	if err != nil {
		return []Kecamatan{}
	}
	kabkotRef := c.serviceKabkot.FindById(kabkotId)
	if (kabkotRef == kabkot_v010.Kabkot{}) {
		return []Kecamatan{}
	}

	return c.service.FindByKabkot(kabkotId)
}

func (c *controllerKecamatan) Create(ctx *gin.Context) (Kecamatan, error) {
	var kecamatan Kecamatan
	var kecamatanCreate KecamatanCreate
	err := ctx.ShouldBindJSON(&kecamatanCreate)
	if err != nil {
		return Kecamatan{}, err
	}

	kabkotRef := c.serviceKabkot.FindById(kecamatanCreate.KabkotID)
	if (kabkotRef == kabkot_v010.Kabkot{}) {
		return Kecamatan{}, errors.New("Kab/Kota tidak valid")
	}

	kecamatan.Name = kecamatanCreate.Name
	kecamatan.Kabkot = kabkotRef
	kecamatan.KabkotID = kabkotRef.ID

	result, err := c.service.Create(kecamatan)
	if err != nil {
		return Kecamatan{}, err
	}
	return result, nil
}

func (c *controllerKecamatan) Update(ctx *gin.Context) error {
	var kecamatan Kecamatan
	var kecamatanEdit KecamatanEdit

	err := ctx.ShouldBindJSON(&kecamatanEdit)
	if err != nil {
		return err
	}

	kecamatan = c.service.FindById(kecamatanEdit.ID)
	if (kecamatan == Kecamatan{}) {
		return errors.New("Kecamatan tidak valid")
	}

	kecamatan.Name = kecamatanEdit.Name
	kecamatan.UpdatedAt = time.Now()
	kabkotRef := c.serviceKabkot.FindById(kecamatanEdit.KabkotID)
	kecamatan.Kabkot = kabkotRef
	kecamatan.KabkotID = kabkotRef.ID

	err = c.service.Update(kecamatan)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerKecamatan) Delete(ctx *gin.Context) error {
	var kecamatan Kecamatan
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	kecamatan = c.service.FindById(id)
	if (kecamatan == Kecamatan{}) {
		return errors.New("Kecamatan tidak valid")
	}

	err = c.service.Delete(kecamatan)
	if err != nil {
		return err
	}
	return nil
}
