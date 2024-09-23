package kabkot_v010

import (
	"errors"
	"rsudlampung/modules/provinsi_v010"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type KabkotController interface {
	FindAll() []Kabkot
	FindByProvinsi(ctx *gin.Context) []Kabkot
	Create(ctx *gin.Context) (Kabkot, error)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ChangeStatus(ctx *gin.Context) error
}

type controllerKabkot struct {
	service         KabkotService
	serviceProvinsi provinsi_v010.ProvinsiService
}

func NewKabkotController(db *gorm.DB) KabkotController {

	return &controllerKabkot{
		service:         NewKabkotService(db),
		serviceProvinsi: provinsi_v010.NewProvinsiService(db),
	}
}

func (c *controllerKabkot) FindAll() []Kabkot {
	return c.service.FindAll()
}

func (c *controllerKabkot) FindByProvinsi(ctx *gin.Context) []Kabkot {
	provinsiId, err := strconv.ParseUint(ctx.Param("provinsi_id"), 10, 64)
	if err != nil {
		return []Kabkot{}
	}
	provinsiRef := c.serviceProvinsi.FindById(provinsiId)
	if (provinsiRef == provinsi_v010.Provinsi{}) {
		return []Kabkot{}
	}

	return c.service.FindByProvinsi(provinsiId)
}

func (c *controllerKabkot) Create(ctx *gin.Context) (Kabkot, error) {
	var kabkot Kabkot
	var kabkotCreate KabkotCreate
	err := ctx.ShouldBindJSON(&kabkotCreate)
	if err != nil {
		return Kabkot{}, err
	}

	provinsiRef := c.serviceProvinsi.FindById(kabkotCreate.ProvinsiID)
	kabkot.Name = kabkotCreate.Name
	kabkot.Provinsi = provinsiRef
	kabkot.ProvinsiID = provinsiRef.ID

	result, err := c.service.Create(kabkot)
	if err != nil {
		return Kabkot{}, err
	}
	return result, nil
}

func (c *controllerKabkot) Update(ctx *gin.Context) error {
	var kabkot Kabkot
	var kabkotEdit KabkotEdit

	err := ctx.ShouldBindJSON(&kabkotEdit)
	if err != nil {
		return err
	}

	kabkot = c.service.FindById(kabkotEdit.ID)
	if (kabkot == Kabkot{}) {
		return errors.New("Kab/Kota tidak valid")
	}

	kabkot.Name = kabkotEdit.Name
	kabkot.UpdatedAt = time.Now()
	provinsiRef := c.serviceProvinsi.FindById(kabkotEdit.ProvinsiID)
	kabkot.Provinsi = provinsiRef
	kabkot.ProvinsiID = provinsiRef.ID
	kabkot.Fasprof = kabkotEdit.Fasprof
	kabkot.IN = kabkotEdit.IN
	kabkot.Status = kabkotEdit.Status

	err = c.service.Update(kabkot)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerKabkot) Delete(ctx *gin.Context) error {
	var kabkot Kabkot
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	kabkot = c.service.FindById(id)
	if (kabkot == Kabkot{}) {
		return errors.New("Kabupaten/Kota tidak valid")
	}

	err = c.service.Delete(kabkot)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerKabkot) ChangeStatus(ctx *gin.Context) error {
	var kabkot Kabkot
	id, err1 := strconv.ParseUint(ctx.Param("id"), 10, 64)
	status, err2 := strconv.ParseBool(ctx.Param("status"))

	if err1 != nil || err2 != nil {
		return errors.New("Data tidak valid")
	}

	kabkot = c.service.FindById(id)
	if (kabkot == Kabkot{}) {
		return errors.New("Data tidak valid")
	}
	kabkot.Status = status

	err := c.service.Update(kabkot)
	if err != nil {
		return err
	}
	return nil
}
