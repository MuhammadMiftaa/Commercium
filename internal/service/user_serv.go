package service

import (
	"errors"

	"commercium/internal/entity"
	"commercium/internal/helper"
	"commercium/internal/repository"
)

type UsersService interface {
	Register(user *entity.UsersRequest) (entity.Users, error)
	Login(user *entity.UsersRequest) (*string, error)
	GetAllUsers() ([]entity.Users, error)
	GetUserByID(id int) (entity.Users, error)
	GetUserByEmail(email string) (entity.Users, error)
	CreateUser(user entity.UsersRequest) (entity.Users, error)
	UpdateUser(id int, userNew entity.UsersRequest) (entity.Users, error)
	DeleteUser(id int) (entity.Users, error)
}

type usersService struct {
	userRepository repository.UsersRepository
}

func NewUsersService(usersRepository repository.UsersRepository) UsersService {
	return &usersService{usersRepository}
}

func (user_serv *usersService) Register(user *entity.UsersRequest) (entity.Users, error) {
	// VALIDASI APAKAH USERNAME, FULLNAME, EMAIL, PASSWORD KOSONG
	if user.Username == "" || user.Fullname == "" || user.Email == "" || user.Password == "" {
		return entity.Users{}, errors.New("username, fullname, email, and password cannot be blank")
	}

	// VALIDASI UNTUK USERNAME AGAR TIDAK BERISI SPASI DAN HANYA MENGANDUNG ALFABET DAN NUMERIK
	if isValid := helper.UsernameValidator(user.Username); !isValid {
		return entity.Users{}, errors.New("usernames can only contain letters and numbers, with no spaces allowed")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := helper.EmailValidator(user.Email); !isValid {
		return entity.Users{}, errors.New("please enter a valid email address")
	}

	// MENGECEK APAKAH USERNAME DAN EMAIL SUDAH DIGUNAKAN
	userExist, err := user_serv.userRepository.GetUserByEmail(user.Email)
	if err == nil && (userExist.Username != "" || userExist.Email != "") {
		return entity.Users{}, errors.New("username or email already exists")
	}

	// VALIDASI PASSWORD SUDAH SESUAI, MIN 8 KARAKTER, MENGANDUNG ALFABET DAN NUMERIK
	hasMinLen, hasLetter, hasDigit := helper.PasswordValidator(user.Password)
	if !hasMinLen {
		return entity.Users{}, errors.New("password must be at least 8 characters long")
	}
	if !hasLetter {
		return entity.Users{}, errors.New("password must contain at least one letter")
	}
	if !hasDigit {
		return entity.Users{}, errors.New("password must contain at least one number")
	}

	// HASHING PASSWORD MENGGUNAKAN BCRYPT
	hashedPassword, err := helper.PasswordHashing(user.Password)
	if err != nil {
		return entity.Users{}, err
	}
	user.Password = hashedPassword

	//  MENGUBAH TIPE USER REQUEST KE ENTITY USER
	newUser := entity.Users{
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
		Role:     "user",
	}

	return user_serv.userRepository.CreateUser(newUser)
}

func (user_serv *usersService) Login(user *entity.UsersRequest) (*string, error) {
	userExist, err := user_serv.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !helper.ComparePass(userExist.Password, user.Password) {
		return nil, errors.New("password is incorrect")
	}

	token, err := helper.GenerateToken(userExist.Username, userExist.Email)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (user_serv *usersService) GetAllUsers() ([]entity.Users, error) {
	return user_serv.userRepository.GetAllUsers()
}

func (user_serv *usersService) GetUserByID(id int) (entity.Users, error) {
	return user_serv.userRepository.GetUserByID(id)
}

func (user_serv *usersService) GetUserByEmail(email string) (entity.Users, error) {
	return user_serv.userRepository.GetUserByEmail(email)
}

func (user_serv *usersService) CreateUser(userRequest entity.UsersRequest) (entity.Users, error) {
	// VALIDASI APAKAH USERNAME, FULLNAME, EMAIL, PASSWORD KOSONG
	if userRequest.Username == "" || userRequest.Fullname == "" || userRequest.Email == "" || userRequest.Password == "" {
		return entity.Users{}, errors.New("username, fullname, email, and password cannot be blank")
	}

	// VALIDASI UNTUK USERNAME AGAR TIDAK BERISI SPASI DAN HANYA MENGANDUNG ALFABET DAN NUMERIK
	if isValid := helper.UsernameValidator(userRequest.Username); !isValid {
		return entity.Users{}, errors.New("usernames can only contain letters and numbers, with no spaces allowed")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := helper.EmailValidator(userRequest.Email); !isValid {
		return entity.Users{}, errors.New("please enter a valid email address")
	}

	// MENGECEK APAKAH USERNAME SUDAH DIGUNAKAN
	if _, err := user_serv.userRepository.GetUserByEmail(userRequest.Email); err == nil {
		return entity.Users{}, errors.New("username already exists")
	}

	// VALIDASI PASSWORD SUDAH SESUAI, MIN 8 KARAKTER, MENGANDUNG ALFABET DAN NUMERIK
	hasMinLen, hasLetter, hasDigit := helper.PasswordValidator(userRequest.Password)
	if !hasMinLen {
		return entity.Users{}, errors.New("password must be at least 8 characters long")
	}
	if !hasLetter {
		return entity.Users{}, errors.New("password must contain at least one letter")
	}
	if !hasDigit {
		return entity.Users{}, errors.New("password must contain at least one number")
	}

	// HASHING PASSWORD MENGGUNAKAN BCRYPT
	hashedPassword, err := helper.PasswordHashing(userRequest.Password)
	if err != nil {
		return entity.Users{}, err
	}
	userRequest.Password = hashedPassword

	//  MENGUBAH TIPE USER REQUEST KE ENTITY USER
	user := entity.Users{
		Username: userRequest.Username,
		Fullname: userRequest.Fullname,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Role:     "user",
	}

	return user_serv.userRepository.CreateUser(user)
}

func (user_serv *usersService) UpdateUser(id int, userNew entity.UsersRequest) (entity.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI UPDATE
	user, err := user_serv.userRepository.GetUserByID(id)
	if err != nil {
		return entity.Users{}, err
	}

	// VALIDASI APAKAH FULLNAME & EMAIL KOSONG
	if userNew.Fullname == "" && userNew.Email == "" {
		return entity.Users{}, errors.New("fullname and email cannot be blank")
	}

	// VALIDASI APAKAH FULLNAME / EMAIL SUDAH DI INPUT
	if userNew.Fullname != "" {
		user.Fullname = userNew.Fullname
	}

	if userNew.Email != "" {
		// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
		if isValid := helper.EmailValidator(userNew.Email); !isValid {
			return entity.Users{}, errors.New("please enter a valid email address")
		}
		// MENGECEK APAKAH EMAIL SUDAH DIGUNAKAN
		existingUser, _ := user_serv.userRepository.GetUserByEmail(userNew.Email)
		if existingUser.ID != 0 && existingUser.ID != user.ID {
			return entity.Users{}, errors.New("email already in use by another user")
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
