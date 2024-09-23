package group_v010

import (
	"log"
	"rsudlampung/helper"
	"time"

	"gorm.io/gorm"
)

type GroupService interface {
	Create(Group) (Group, error)
	Update(Group) error
	Delete(Group) error
	FindAll() []Group
	FindById(groupId uint64) Group
}

type groupService struct {
	conn *gorm.DB
}

func NewGroupService(db *gorm.DB) GroupService {
	configEnv, errEnv := helper.LoadConfig("../.")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}
	am := configEnv.AutoMigrate

	if am == "on" {
		db.AutoMigrate(&Group{})
	}

	return &groupService{
		conn: db,
	}
}

func (service *groupService) Create(group Group) (Group, error) {

	group.UpdatedAt = time.Now()
	err := service.conn.Create(&group).Error
	if err != nil {
		return Group{}, err
	}
	return group, nil
}

func (service *groupService) Update(group Group) error {
	group.CreatedAt = time.Now()
	err := service.conn.Save(&group).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *groupService) Delete(group Group) error {
	err := service.conn.Delete(&group).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *groupService) FindAll() []Group {
	var groups []Group
	service.conn.Find(&groups)
	return groups
}

func (service *groupService) FindById(groupId uint64) Group {
	var group Group
	service.conn.Where("id=?", groupId).Find(&group)
	return group
}
