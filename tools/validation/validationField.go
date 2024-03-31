package validation

import (
	"errors"
	"github.com/google/uuid"
	"slices"
	"unicode"

	"regexp"
	"strings"
)

var paramInvalid = " uuid is invalid"

func ValidatePassword(password, fieldName string) error {
	var charList = []rune{'@', '$', '_', '.', '-', '+', '#', '!', '*', '?', '&'}

	if password == "" {
		return errors.New("password cannot be blank")
	}
	var ErrInvalidPassword = errors.New(`the password should at least have 7 letters, at least 1 number, at least 1 upper case, at least 1 special character(@,$,_,.,-,+,#,!,*,?,&)`)
	if len(password) >= 30 && len(password) <= 8 {
		return ErrInvalidPassword
	}
	var num, lower, upper, spec bool
	for _, pass := range password {
		switch {
		case unicode.IsDigit(pass):
			num = true
		case unicode.IsUpper(pass):
			upper = true
		case unicode.IsLower(pass):
			lower = true
		case slices.Contains(charList, pass):
			spec = true
		}
	}
	if num && lower && upper && spec {
		return nil
	}
	return ErrInvalidPassword
}

func ValidationLogin(name, fieldName string) error {
	if strings.TrimSpace(name) != "" {
		if !regexp.MustCompile("^[^~@$%^&*+=`|{}:;!?\\\"()\\[\\]-]+(?:.*\\d)?(?:.*[a-z])?(?:.*[A-Z])?(?:.*[-,./ #])?(?:[ ])?$").MatchString(name) {
			return errors.New(fieldName + " must be only letters and numbers and" +
				" special characters: (.,-/ #)")
		}
	}
	return nil
}

func ValidationEmail(email, fieldName string) error {
	if email != "" {
		email = strings.TrimSpace(email)
		if !regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(email) {
			return errors.New("email must be like this format: [username]@[domain" +
				"].[domain]")
		}
	}
	return nil
}

func UUIDValidation(uuidValue string) error {
	//if regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$").MatchString(uuidValue){
	//}
	_, err := uuid.Parse(uuidValue)
	if err != nil {
		return errors.New(paramInvalid)
	}
	return nil
}

func ValidatorHexColor(hexColor, fieldName string) error {
	if hexColor != "" {
		match, err := regexp.MatchString("^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$", hexColor)
		if err != nil {
			return err
		}
		if !match {
			return errors.New("invalid hex color format: " + fieldName)
		}
	}
	return nil
}

func ValidatorNumber(number, fieldName string) error {
	if number != "" {
		match, err := regexp.MatchString("^[0-9]*$", string(number))
		if err != nil {
			return err
		}
		if !match {
			return errors.New(fieldName + "invalid number format: ")
		}
	}
	return nil
}

func ValidatorPhone(fieldName, phone string) error {
	if phone != "" {

		r := regexp.MustCompile(`^\+998[0-9]{2}[0-9]{7}$`)
		result := r.MatchString(phone)
		if !result {
			return errors.New("phone number must be like this format: +YYYZZXXXXXXX")
		}
	}
	return nil
}

func ValidatorMonth(value int64) error {
	if value < 0 || value > 12 {
		return errors.New("months must be from 1 to 12")
	}
	return nil
}
