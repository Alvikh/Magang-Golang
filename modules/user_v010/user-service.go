package user_v010

import (
	"log"
	"rsudlampung/helper"
	"time"

	"gorm.io/gorm"
)

type UserService interface {
	Create(User) (User, error)
	Update(User) error
	Delete(User) error
	FindAll() []User
	FindByNik(userNik string) User
}

type userService struct {
	conn *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	configEnv, errEnv := helper.LoadConfig("../.")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}
	am := configEnv.AutoMigrate

	if am == "on" {
		db.AutoMigrate(&User{})
	}

	return &userService{
		conn: db,
	}
}

func (service *userService) Create(user User) (User, error) {

	user.UpdatedAt = time.Now()
	err := service.conn.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (service *userService) Update(user User) error {
	user.CreatedAt = time.Now()
	err := service.conn.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) Delete(user User) error {
	err := service.conn.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) FindAll() []User {
	var users []User
	service.conn.Find(&users)
	return users
}

func (service *userService) FindByNik(userNik string) User {
	var user User
	service.conn.Where("nik=?", userNik).Find(&user)
	return user
}
