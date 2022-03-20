package userInput

import (
	"fmt"
)

func TakeUserInput(firstName string, lastName string, email string, numberOfTickets int)  {
	fmt.Printf("Please provide your necessary details to book your ticket.\n")
	fmt.Print("First Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Number of Tickets: ")
	fmt.Scan(&numberOfTickets)
}