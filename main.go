package main

import (
	"Booking-App/model"
	"Booking-App/printBookings"
	"Booking-App/userInput"
	"Booking-App/welcome"
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


  //Array of UserData
  var bookings = make([]model.User, 0)

  //Welcome the user
  welcome.WelcomeMsgPrint("GoLang")

  //Take user input and store in variables inside main.
  for i := 0; i < 2; i++ {
    firstName, lastName, email, numberOfTickets, venue := userInput.TakeUserInput(venues)
    var userData = model.User {
      FirstName: firstName, 
      LastName: lastName, 
      Email: email,
      NumberOfTickets: numberOfTickets,
      Venue: venue,
    }
    bookings = append(bookings, userData)
  }

  //Print the bookings to the console
  printBookings.PrintBookings(bookings)
}