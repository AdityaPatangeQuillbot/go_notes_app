package core

import (
	"errors"
	"main/models"
	"regexp"
)

var ErrInvalidEmail = errors.New("invalid email")
var ErrPasswordCannotBeEmpty = errors.New("password cannot be empty")

func IsEmailValid(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)

	if !match || email == "" {
		return false
	}

	return true
}

func ValidateSignupRequest(signupRequest models.SignupRequest) (models.SignupRequest, error) {
	email := signupRequest.Username

	if !IsEmailValid(email) {
		return models.SignupRequest{}, ErrInvalidEmail
	}

	if signupRequest.Password == "" || signupRequest.ConfirmPassword == "" {
		return models.SignupRequest{}, ErrPasswordCannotBeEmpty
	}

	return signupRequest, nil
}
