package service

import (
	"errors"

	"commercium/internal/entity"
	"commercium/internal/helper"
	"commercium/internal/repository"
)

type UsersService interface {
	GetAllUsers() ([]entity.Users, error)
	GetUserByID(id int) (entity.Users, error)
	GetUserByUsername(username string) (entity.Users, error)
	CreateUser(user entity.Users) (entity.Users, error)
	UpdateUser(id int, userNew entity.Users) (entity.Users, error)
	DeleteUser(id int) (entity.Users, error)
}

type usersService struct {
	userRepository repository.UsersRepository
}

func NewUsersService(usersRepository repository.UsersRepository) *usersService {
	return &usersService{usersRepository}
}

func (user_serv *usersService) GetAllUsers() ([]entity.Users, error) {
	return user_serv.userRepository.GetAllUsers()
}

func (user_serv *usersService) GetUserByID(id int) (entity.Users, error) {
	return user_serv.userRepository.GetUserByID(id)
}

func (user_serv *usersService) GetUserByUsername(username string) (entity.Users, error) {
	return user_serv.userRepository.GetUserByUsername(username)
}

func (user_serv *usersService) CreateUser(user entity.Users) (entity.Users, error) {
	// VALIDASI UNTUK USERNAME AGAR TIDAK BERISI SPASI DAN HANYA MENGANDUNG ALFABET DAN NUMERIK
	if isValid := helper.UsernameValidator(user.Username); !isValid {
		return entity.Users{}, errors.New("usernames can only contain letters and numbers, with no spaces allowed")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := helper.EmailValidator(user.Email); !isValid {
		return entity.Users{}, errors.New("please enter a valid email address")
	}

	// MENGISI ROLE DEFAULT USER
	if user.Role == "" {
		user.Role = "user"
	}

	// VALIDASI PASSWORD SUDAH SESUAI, MIN 8 KARAKTER, MENGANDUNG ALFABET DAN NUMERIK
	hasMinLen, hasLetter, hasDigit := helper.PasswordValidator(user.Password)
	if !hasMinLen {
		return user, errors.New("password must be at least 8 characters long")
	}
	if !hasLetter {
		return user, errors.New("password must contain at least one letter")
	}
	if !hasDigit {
		return user, errors.New("password must contain at least one number")
	}

	// HASHING PASSWORD MENGGUNAKAN BCRYPT
	hashedPassword, err := helper.PasswordHashing(user.Password)
	if err != nil {
		return entity.Users{}, err
	}
	user.Password = hashedPassword

	return user_serv.userRepository.CreateUser(user)
}

func (user_serv *usersService) UpdateUser(id int, userNew entity.Users) (entity.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI UPDATE
	user, err := user_serv.userRepository.GetUserByID(id)
	if err != nil {
		return entity.Users{}, err
	}

	// VALIDASI APAKAH FULLNAME / EMAIL SUDAH DI INPUT
	if userNew.Fullname != "" {
		user.Fullname = userNew.Fullname
	}
	if userNew.Email != "" {
		if isValid := helper.EmailValidator(userNew.Email); !isValid {
			return entity.Users{}, errors.New("please enter a valid email address")
		}
		user.Email = userNew.Email
	}

	return user_serv.userRepository.UpdateUser(user)
}

func (user_serv *usersService) DeleteUser(id int) (entity.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI UPDATE
	user, err := user_serv.userRepository.GetUserByID(id)
	if err != nil {
		return entity.Users{}, err
	}

	return user_serv.userRepository.DeleteUser(user)
}
