package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		ID   uuid.UUID
		Name string
	}

	UserRepository struct {
		Db *gorm.DB
	}
)

func (u User) TableName() string {
	return "users"
}

func (u UserRepository) GetUsers() (users []User, err error) {
	users = []User{}
	err = u.Db.Find(&users).Error
	return users, err
}

func (u UserRepository) GetUserById(id uuid.UUID) (user User, err error) {
	err = u.Db.Model(&user).Where("id", id).First(&user).Error
	return user, err
}
