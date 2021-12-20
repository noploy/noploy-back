package repositories

import (
	"github.com/noploy/noploy-back/repositories/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(entity *entities.UserEntity) error
	Delete(id int64) error
	FindById(id int64) (*entities.UserEntity, error)
	FindAll() ([]*entities.UserEntity, error)
}

type UserRepositoryContext struct {
	DB *gorm.DB
}

func (u *UserRepositoryContext) Create(entity *entities.UserEntity) error {
	err := u.DB.Create(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryContext) Delete(id int64) error {
	err := u.DB.Where("id = ?", id).Delete(&entities.UserEntity{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryContext) FindById(id int64) (*entities.UserEntity, error) {
	var user entities.UserEntity
	err := u.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepositoryContext) FindAll() ([]*entities.UserEntity, error) {
	var users []*entities.UserEntity
	err := u.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryContext{DB: db}
}
