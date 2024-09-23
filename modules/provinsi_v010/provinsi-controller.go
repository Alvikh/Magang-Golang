package provinsi_v010

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProvinsiController interface {
	FindAll() []Provinsi
	Create(ctx *gin.Context) (Provinsi, error)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controllerProvinsi struct {
	service ProvinsiService
}

func NewProvinsiController(db *gorm.DB) ProvinsiController {

	return &controllerProvinsi{
		service: NewProvinsiService(db),
	}
}

func (c *controllerProvinsi) FindAll() []Provinsi {
	return c.service.FindAll()
}

func (c *controllerProvinsi) Create(ctx *gin.Context) (Provinsi, error) {
	var provinsi Provinsi
	err := ctx.ShouldBindJSON(&provinsi)
	if err != nil {
		return Provinsi{}, err
	}
	result, err := c.service.Create(provinsi)
	if err != nil {
		return Provinsi{}, err
	}
	return result, nil
}

func (c *controllerProvinsi) Update(ctx *gin.Context) error {
	var provinsi Provinsi
	err := ctx.ShouldBindJSON(&provinsi)
	if err != nil {
		return err
	}

	provinsiRef := c.service.FindById(provinsi.ID)
	provinsi.CreatedAt = provinsiRef.CreatedAt
	err = c.service.Update(provinsi)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerProvinsi) Delete(ctx *gin.Context) error {
	var provinsi Provinsi
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	provinsi = c.service.FindById(id)
	if (provinsi == Provinsi{}) {
		return errors.New("Provinsi tidak valid")
	}

	err = c.service.Delete(provinsi)
	if err != nil {
		return err
	}
	return nil
}
