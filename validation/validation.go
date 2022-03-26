package validation

import (
	"Booking-App/print"
	"strconv"
	"unicode"
)

func ValidateString(str string) bool {
	for _, s := range str {
		if unicode.IsSpace(s) {
			print.PrintError("Invalid Input")
			return false
		}
	}
	return true
}

func ValidateUsername(str string) bool {
	var validString = ValidateString(str)

	if !validString {
		return false
	}

	var runes = []rune(str)

	if unicode.IsNumber(runes[0]) {
		print.PrintError("Username must start with a letter")
		return false
	}

	return true
}

func ValidateNumber(str string) bool {
	var number, err = strconv.ParseUint(str, 10, 64)
	if err != nil || int(number) < 1 {
		print.PrintError("Invalid Input")
		return false
	} else {
		return true
	}
}

func ValidatePassword(password string) bool {

	var validString = ValidateString(password)

	if !validString {
		return false
	}

	var containsUpper = false
	var containsNumber = false
	var containsLowerCase = false

	if len(password) < 8 {
		print.PrintError("Password must be at least 8 characters long")
		return false
	}

	for _, s := range password {
		if unicode.IsUpper(s) {
			containsUpper = true
		} else if unicode.IsNumber(s) {
			containsNumber = true
		} else if unicode.IsLower(s) {
			containsLowerCase = true
		}
	}

	if !containsUpper || !containsNumber || !containsLowerCase {
		print.PrintError("Password must contain at least one uppercase letter, one number and one lowercase letter")
		return false
	}

	return true
}

func ValidateEmail(email string) bool {

	var validString = ValidateString(email)

	if !validString {
		return false
	}

	var containsAt = false
	var containsDot = false

	for _, s := range email {
		if s == '@' {
			containsAt = true
		} else if s == '.' {
			containsDot = true
		}
	}

	if !containsAt || !containsDot {
		print.PrintError("Email must have a valid format")
		return false
	}

	return true
}
