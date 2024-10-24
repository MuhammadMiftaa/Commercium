package helper

import (
	"regexp"
	"unicode"

	"commercium/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

func UsernameValidator(str string) bool {
	username_validator := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return username_validator.MatchString(str)
}

func EmailValidator(str string) bool {
	email_validator := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return email_validator.MatchString(str)
}

func PasswordValidator(str string) (bool, bool, bool) {
	var hasLetter, hasDigit, hasMinLen bool
	for _, char := range str {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if len(str) >= 8 {
		hasMinLen = true
	}

	return hasMinLen, hasLetter, hasDigit
}

func PasswordHashing(str string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func ConvertToResponseType(data interface{}) interface{} {
	switch v := data.(type) {
	case entity.Users:
		return entity.UsersResponse{
			ID:       v.ID,
			Username: v.Username,
			Fullname: v.Fullname,
			Email:    v.Email,
			Role:     v.Role,
		}
	case entity.Products:
		return entity.ProductsResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
		}
	case entity.Orders:
		return entity.OrdersResponse{
			ID:         v.ID,
			ProductID:  v.ProductID,
			UserID:     v.UserID,
			Quantity:   v.Quantity,
			TotalPrice: v.TotalPrice,
			Status:     v.Status,
		}
	default:
		return nil
	}
}
