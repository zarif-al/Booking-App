package main

import (
	"Booking-App/userInput"
	"Booking-App/welcome"
	"fmt"
)

func main(){
  const totalTickets = 50
  
  //Map of venues, using a map instead of array to try out maps.
  var venues = map[int]string{
    1: "Vienna",
    2: "Dhaka",
    3: "Tokyo",
    4: "New York",
    5: "London",
   }

  //User Data Struct
  type UserData struct {
    firstName string
    lastName string
    email string
    numberOfTickets uint
    venue string
  }

  //Array of UserData
  var bookings = make([]UserData, 0)

  //Welcome the user
  welcome.WelcomeMsgPrint("GoLang")

  //Take user input and store in variables inside main.
  for i := 0; i < 5; i++ {
    firstName, lastName, email, numberOfTickets, venue := userInput.TakeUserInput(venues)
    var userData = UserData {
      firstName: firstName, 
      lastName: lastName, 
      email: email,
      numberOfTickets: numberOfTickets,
      venue: venue,
    }
    bookings = append(bookings, userData)
  }

  //Print the bookings to the console
  fmt.Printf("List of bookings %v", bookings)
}