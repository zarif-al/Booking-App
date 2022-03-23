package userInput

import (
	"fmt"
)

func TakeUserInput(venues map[int]string)(string, string, string, uint, string)  {
	var firstName string
  var lastName string
  var email string
  var numberOfTickets uint
	var venue int

	fmt.Println("First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Email: ")
	fmt.Scan(&email)
	fmt.Println("Number of Tickets: ")
	fmt.Scan(&numberOfTickets)
	fmt.Println("Please select a venue (number):")
	for key, value := range venues {
		fmt.Println(key, value)
	}
	fmt.Println("Venue: ")
	fmt.Scan(&venue)
	fmt.Println("Thank You For Booking Your Tickets! See You Soon!")

	return firstName, lastName, email, numberOfTickets, venues[venue]
}