package main

import (
	"Booking-App/userInput"
	"Booking-App/welcome"
	"fmt"
)

func main(){
  var firstName string
  var lastName string
  var email string
  var numberOfTickets int

  welcome.WelcomeMsgPrint("GoLang")
  userInput.TakeUserInput(firstName, lastName, email, numberOfTickets)
  fmt.Print(firstName + " " + lastName + " " + email + " " + numberOfTickets)
}