package repositories

import (
	"errors"

	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindOne(string) (*domain.User, error)
	FindOneByUsername(string) (*domain.User, error)
	Save(domain.User) error
	Delete(string) error
	Automigrate() error
}

type SQLUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	var userRepository UserRepository = SQLUserRepository{db}
	return userRepository
}

var _ UserRepository = &SQLUserRepository{}

func (r SQLUserRepository) Save(user domain.User) error {
	var err error
	var prev UserEntity
	err = r.db.First(&prev, "id = ?", user.ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = r.db.Create(UserEntityFromDomain(user)).Error
		return err
	}

	err = r.db.Save(UserEntityFromDomain(user)).Error
	return err
}

func (r SQLUserRepository) FindOne(id string) (*domain.User, error) {
	var user UserEntity
	err := r.db.Where("id = ?", id).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = &domain.UserNotFound{ID: id}
	}
	if err != nil {
		return nil, err
	}
	userDomain := UserEntityToDomain(user)
	return &userDomain, nil
}

func (r SQLUserRepository) FindOneByUsername(username string) (*domain.User, error) {
	var user UserEntity
	err := r.db.Where("username = ?", username).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = &domain.UsernameNotFound{Username: username}
	}
	if err != nil {
		return nil, err
	}
	userDomain := UserEntityToDomain(user)
	return &userDomain, nil
}

func (r SQLUserRepository) Delete(id string) error {
	var user UserEntity
	err := r.db.Where("id = ?", id).Delete(&user).Error
	return err
}

func (r SQLUserRepository) Automigrate() error {
	return r.db.AutoMigrate(&UserEntity{})
}
