package auth

import (
	"Booking-App/model"
	"Booking-App/validation"
	"bufio"
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

func InputValidString(scanner *bufio.Scanner, msg string) string {
	fmt.Print(msg)
	scanner.Scan()
	for !validation.ValidateString(scanner.Text()) {
		fmt.Print(msg)
		scanner.Scan()
	}
	return scanner.Text()
}

func Login(bookings []model.Booking) (model.Booking, bool) {
	var userName string
	var password string
	scanner := bufio.NewScanner(os.Stdin)
	userName = InputValidString(scanner, color.InCyan("UserName: "))
	password = InputValidString(scanner, color.InCyan("Password: "))
	for _, user := range bookings {
		if user.UserName == userName && user.Password == password {
			return user, true
		}
	}
	return model.Booking{}, false
}
