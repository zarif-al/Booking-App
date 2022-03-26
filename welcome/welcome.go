package welcome

import (
	"Booking-App/print"
	"Booking-App/validation"
	"fmt"
	"strconv"

	"bufio"
	"os"

	"github.com/TwiN/go-color"
)

func InputOption(scanner *bufio.Scanner) int {
	scanner.Scan()
	for !validation.ValidateNumber(scanner.Text()) {
		fmt.Print(color.InBlue("Please select an option: "))
		scanner.Scan()
	}
	option, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if option < 1 || option > 6 {
		print.PrintError("Invalid option. Please select an option between 1 and 6.")
		return InputOption(scanner)
	}

	return int(option)
}

func WelcomeSection(conferenceName string) int {
	var option int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to", color.InBold(conferenceName), "conference ticket booking application.")
	fmt.Println("What would you like to do?")
	fmt.Println(color.InCyan("1."), "Book Tickets")
	fmt.Println(color.InCyan("2."), "Change Number of Tickets")
	fmt.Println(color.InCyan("3."), "Change Venue")
	fmt.Println(color.InCyan("4."), "Print Your Booking Details")
	fmt.Println(color.InCyan("5."), "Cancel Booking")
	fmt.Println(color.InCyan("6."), "Exit")
	fmt.Print(color.InBlue("Please select an option: "))
	option = InputOption(scanner)
	return option
}
