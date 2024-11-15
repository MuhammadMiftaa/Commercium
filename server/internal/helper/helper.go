package helper

import (
	"errors"
	"regexp"
	"time"
	"unicode"

	"commercium/internal/entity"

	"github.com/dgrijalva/jwt-go"
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
			Category:    v.Category,
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

var secretKey = "pojq09720ef1ko0f1h9iego2010j20240"

func GenerateToken(username string, email string, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"role":     role,
		"exp":      expirationTime.Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(cookie string) (interface{}, error) {
	token, _ := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("sign in to preceed")
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New("sign in to preceed")
	}

	return token.Claims.(jwt.MapClaims), nil
}

func ComparePass(hashPassword, reqPassword string) bool {
	hash, pass := []byte(hashPassword), []byte(reqPassword)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		return false
	}

	return true
}
