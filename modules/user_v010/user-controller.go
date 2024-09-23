package user_v010

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	FindAll() []User
	FindByNik(ctx *gin.Context) (User, error)
	Create(ctx *gin.Context) (User, error)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controllerUser struct {
	service UserService
}

func NewUserController(db *gorm.DB) UserController {

	return &controllerUser{
		service: NewUserService(db),
	}
}

func (c *controllerUser) FindAll() []User {
	return c.service.FindAll()
}

func (c *controllerUser) FindByNik(ctx *gin.Context) (User, error) {
	var user User
	nik := ctx.Param("nik")
	user = c.service.FindByNik(nik)
	if (user == User{}) {
		return User{}, errors.New("User tidak valid")
	}
	return user, nil
}

func (c *controllerUser) Create(ctx *gin.Context) (User, error) {
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return User{}, err
	}
	result, err := c.service.Create(user)
	if err != nil {
		return User{}, err
	}
	return result, nil
}

func (c *controllerUser) Update(ctx *gin.Context) error {
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	err = c.service.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerUser) Delete(ctx *gin.Context) error {
	var user User
	nik := ctx.Param("nik")

	user = c.service.FindByNik(nik)
	if (user == User{}) {
		return errors.New("User tidak valid")
	}

	user.NIK = nik
	err := c.service.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
