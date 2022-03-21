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
	fmt.Print("First Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Number of Tickets: ")
	fmt.Scan(&numberOfTickets)
	fmt.Println("Please select a venue (number):")
	for key, value := range venues {
		fmt.Printf("%d. %s\n", key, value)
	}
	fmt.Print("Venue: ")
	fmt.Scan(&venue)

	return firstName, lastName, email, numberOfTickets, venues[venue]
}