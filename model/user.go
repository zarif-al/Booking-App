package model

import (
	"Booking-App/print"
	"Booking-App/validation"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/TwiN/go-color"
)

//Booking  Struct
type Booking struct {
	UserName        string
	Password        string
	Email           string
	NumberOfTickets uint
	Venue           string
}

func InputUsername(scanner *bufio.Scanner) string {
	print.PrintInputMsg("UserName (No Spaces): ")
	scanner.Scan()
	for !validation.ValidateUsername(scanner.Text()) {
		print.PrintInputMsg("UserName (No Spaces): ")
		scanner.Scan()
	}
	return scanner.Text()
}

func InputPassword(scanner *bufio.Scanner) string {
	print.PrintInputMsg("Password: ")
	scanner.Scan()
	for !validation.ValidatePassword(scanner.Text()) {
		print.PrintInputMsg("Password: ")
		scanner.Scan()
	}
	return scanner.Text()
}

func InputEmail(scanner *bufio.Scanner) string {
	print.PrintInputMsg("Email: ")
	scanner.Scan()
	for !validation.ValidateEmail(scanner.Text()) {
		print.PrintInputMsg("Email: ")
		scanner.Scan()
	}
	return scanner.Text()
}

func InputTicketNumber(scanner *bufio.Scanner, remainingTickets int) uint {
	print.PrintInputMsg("Number of tickets: ")
	scanner.Scan()
	for !validation.ValidateNumber(scanner.Text()) {
		print.PrintInputMsg("Number of tickets: ")
		scanner.Scan()
	}
	number, _ := strconv.ParseUint(scanner.Text(), 10, 64)

	if int(number) > remainingTickets {
		warning := "Sorry, we only have " + strconv.Itoa(remainingTickets) + " tickets left! Please try again."
		print.PrintWarning(warning)
		return InputTicketNumber(scanner, remainingTickets)
	}

	return uint(number)

}

func InputVenue(scanner *bufio.Scanner, venueLength int) int {
	print.PrintInputMsg("Venue (number): ")
	scanner.Scan()
	for !validation.ValidateNumber(scanner.Text()) {
		print.PrintInputMsg("Venue (number): ")
		scanner.Scan()
	}
	number, _ := strconv.ParseUint(scanner.Text(), 10, 64)
	if int(number) > venueLength {
		print.PrintWarning("Invalid venue selection please try again!")
		return InputVenue(scanner, venueLength)
	}
	return int(number)
}

func CreateBooking(venues map[int]string, remainingTickets int) Booking {
	var userName string
	var email string
	var numberOfTickets uint
	var venue int
	var password string
	scanner := bufio.NewScanner(os.Stdin)

	userName = InputUsername(scanner)
	password = InputPassword(scanner)
	email = InputEmail(scanner)
	numberOfTickets = InputTicketNumber(scanner, remainingTickets)
	fmt.Println("These are our venues:")
	fmt.Println(venues)
	venue = InputVenue(scanner, len(venues))
	print.PrintSuccess("Thank You For Booking Your Tickets! See You Soon!")

	var newUser = Booking{
		UserName:        userName,
		Password:        password,
		Email:           email,
		NumberOfTickets: numberOfTickets,
		Venue:           venues[venue],
	}

	return newUser
}

func (u Booking) ChangeTicket(remainingTickets int) int {
	var numberOfTickets uint
	var oldNumberOfTickets = u.NumberOfTickets

	scanner := bufio.NewScanner(os.Stdin)
	numberOfTickets = InputTicketNumber(scanner, remainingTickets)
	u.NumberOfTickets = numberOfTickets

	var difference = numberOfTickets - oldNumberOfTickets

	return int(difference)
}

func (u Booking) ChangeVenue(venues map[int]string) {
	var venue int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("These are our venues:")
	fmt.Println(venues)
	venue = InputVenue(scanner, len(venues))
	u.Venue = venues[venue]
}

func (u Booking) PrintBookingDetails() {
	fmt.Println()
	fmt.Println(color.Ize(color.Cyan, "UserName: "), u.UserName)
	fmt.Println(color.Ize(color.Cyan, "Email: "), u.Email)
	fmt.Println(color.Ize(color.Cyan, "Number Of Tickets: "), u.NumberOfTickets)
	fmt.Println(color.Ize(color.Cyan, "Venue: "), u.Venue)
	fmt.Println()
}
