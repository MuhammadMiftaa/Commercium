package repository

import (
	"commercium/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetAllUsers() ([]entity.Users, error)
	GetUserByID(id int) (entity.Users, error)
	GetUserByUsername(username string) (entity.Users, error)
	CreateUser(user entity.Users) (entity.Users, error)
	UpdateUser(user entity.Users) (entity.Users, error)
	DeleteUser(user entity.Users) (entity.Users, error)
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db}
}

func (user_repo *usersRepository) GetAllUsers() ([]entity.Users, error) {
	var users []entity.Users
	err := user_repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (user_repo *usersRepository) GetUserByID(id int) (entity.Users, error) {
	var user entity.Users
	err := user_repo.db.First(&user, id).Error
	if err != nil {
		return entity.Users{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) GetUserByUsername(username string) (entity.Users, error) {
	var user entity.Users
	err := user_repo.db.First(&user, "username = ?", username).Error
	if err != nil {
		return entity.Users{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) CreateUser(user entity.Users) (entity.Users, error) {
	err := user_repo.db.Create(&user).Error
	if err != nil {
		return entity.Users{}, errors.New("failed to create user")
	}

	return user, nil
}

func (user_repo *usersRepository) UpdateUser(user entity.Users) (entity.Users, error) {
	err := user_repo.db.Save(&user).Error
	if err != nil {
		return entity.Users{}, errors.New("failed to update user")
	}

	return user, nil
}

func (user_repo *usersRepository) DeleteUser(user entity.Users) (entity.Users, error) {
	err := user_repo.db.Delete(&user).Error
	if err != nil {
		return entity.Users{}, errors.New("failed to delete user")
	}

	return user, nil
}
