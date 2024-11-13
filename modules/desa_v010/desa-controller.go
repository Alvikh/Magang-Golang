package desa_v010

import (
	"errors"
	"rsudlampung/modules/kecamatan_v010"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DesaController interface {
	FindAll(ctx *gin.Context) []Desa
	FindByKecamatan(ctx *gin.Context) []Desa
	Create(ctx *gin.Context) (Desa, error)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controllerDesa struct {
	service          DesaService
	serviceKecamatan kecamatan_v010.KecamatanService
}

func NewDesaController(db *gorm.DB) DesaController {
	return &controllerDesa{
		service:          NewDesaService(db),
		serviceKecamatan: kecamatan_v010.NewKecamatanService(db),
	}
}

func (c *controllerDesa) FindAll(ctx *gin.Context) []Desa {
	return c.service.FindAll()
}

func (c *controllerDesa) FindByKecamatan(ctx *gin.Context) []Desa {
	kecamatanId, err := strconv.ParseUint(ctx.Param("kecamatan_id"), 10, 64)
	if err != nil {
		return []Desa{}
	}
	kecamatanRef := c.serviceKecamatan.FindById(kecamatanId)
	if (kecamatanRef == kecamatan_v010.Kecamatan{}) {
		return []Desa{}
	}

	return c.service.FindByKecamatan(kecamatanId)
}

func (c *controllerDesa) Create(ctx *gin.Context) (Desa, error) {
	var desa Desa
	var desaCreate DesaCreate
	err := ctx.ShouldBindJSON(&desaCreate)
	if err != nil {
		return Desa{}, err
	}

	kecamatanRef := c.serviceKecamatan.FindById(desaCreate.KecamatanID)
	if (kecamatanRef == kecamatan_v010.Kecamatan{}) {
		return Desa{}, errors.New("kecamatan tidak valid")
	}

	desa.Name = desaCreate.Name
	desa.Kecamatan = kecamatanRef
	desa.KecamatanID = kecamatanRef.ID

	result, err := c.service.Create(desa)
	if err != nil {
		return Desa{}, err
	}
	return result, nil
}

func (c *controllerDesa) Update(ctx *gin.Context) error {
	var desa Desa
	var desaEdit DesaEdit

	err := ctx.ShouldBindJSON(&desaEdit)
	if err != nil {
		return err
	}

	desa = c.service.FindById(desaEdit.ID)
	if (desa == Desa{}) {
		return errors.New("desa tidak valid")
	}

	desa.Name = desaEdit.Name
	desa.UpdatedAt = time.Now()
	kecamatanRef := c.serviceKecamatan.FindById(desaEdit.KecamatanID)
	desa.Kecamatan = kecamatanRef
	desa.KecamatanID = kecamatanRef.ID

	err = c.service.Update(desa)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerDesa) Delete(ctx *gin.Context) error {
	var desa Desa
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	desa = c.service.FindById(id)
	if (desa == Desa{}) {
		return errors.New("desa tidak valid")
	}

	err = c.service.Delete(desa)
	if err != nil {
		return err
	}
	return nil
}
