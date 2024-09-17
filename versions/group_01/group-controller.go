package group_01

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GroupController interface {
	FindAll() []Group
	FindById(ctx *gin.Context) (Group, error)
	Create(ctx *gin.Context) (Group, error)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controllerGroup struct {
	service GroupService
}

func NewGroupController(db *gorm.DB) GroupController {

	return &controllerGroup{
		service: NewGroupService(db),
	}
}

func (c *controllerGroup) FindAll() []Group {
	return c.service.FindAll()
}

func (c *controllerGroup) FindById(ctx *gin.Context) (Group, error) {
	var group Group
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return Group{}, err
	}

	group = c.service.FindById(id)
	if (group == Group{}) {
		return Group{}, errors.New("Group tidak valid")
	}
	return group, nil
}

func (c *controllerGroup) Create(ctx *gin.Context) (Group, error) {
	var group Group
	err := ctx.ShouldBindJSON(&group)
	if err != nil {
		return Group{}, err
	}
	result, err := c.service.Create(group)
	if err != nil {
		return Group{}, err
	}
	return result, nil
}

func (c *controllerGroup) Update(ctx *gin.Context) error {
	var group Group
	err := ctx.ShouldBindJSON(&group)
	if err != nil {
		return err
	}

	err = c.service.Update(group)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerGroup) Delete(ctx *gin.Context) error {
	var group Group
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	group = c.service.FindById(id)
	if (group == Group{}) {
		return errors.New("Group tidak valid")
	}

	group.ID = id
	err = c.service.Delete(group)
	if err != nil {
		return err
	}
	return nil
}
