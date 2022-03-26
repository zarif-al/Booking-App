package main

import (
	auth "Booking-App/authentication"
	"Booking-App/model"
	"Booking-App/print"
	"Booking-App/welcome"
	"sort"
	"strconv"
)

func main() {
	const totalTickets = 50
	var remainingTickets int = 50
	//Map of venues, using a map instead of array to try out maps.
	var venues = map[int]string{
		1: "Vienna",
		2: "Dhaka",
		3: "Tokyo",
		4: "New York",
		5: "London",
	}

	//Array of Booking Bookings
	var bookings = make([]model.Booking, 0)

	var option int

	for option != 6 {
		option = welcome.WelcomeSection("GoLang")
		switch option {
		case 1:
			if remainingTickets > 0 {
				var booking = model.CreateBooking(venues, remainingTickets)
				remainingTickets = remainingTickets - int(booking.NumberOfTickets)
				print.PrintWarning("Remaining tickets: " + strconv.Itoa(remainingTickets))
				bookings = append(bookings, booking)
			} else {
				print.PrintWarning("Sorry, there are no tickets left!")
			}
			break
		case 2:
			var booking, success = auth.Login(bookings)
			if success {
				difference := booking.ChangeTicket(remainingTickets)
				remainingTickets = remainingTickets - difference
				print.PrintWarning("Remaining tickets: " + strconv.Itoa(remainingTickets))
			} else {
				print.PrintError("Invalid username or password!")
			}
			break
		case 3:
			var booking, success = auth.Login(bookings)
			if success {
				booking.ChangeVenue(venues)
			} else {
				print.PrintError("Invalid username or password!")
			}
		case 4:
			var booking, success = auth.Login(bookings)
			if success {
				booking.PrintBookingDetails()
			} else {
				print.PrintError("Invalid username or password!")
			}
		case 5:
			var booking, success = auth.Login(bookings)
			if success {
				index := sort.Search(len(bookings), func(i int) bool {
					return bookings[i].UserName == booking.UserName
				})

				bookings[index] = bookings[len(bookings)-1]
				bookings = bookings[:len(bookings)-1]
			} else {
				print.PrintError("Invalid username or password!")
			}
		}
	}

}

/*
	Slower method to remove from array but maintains order
   bookings = append(bookings[:index], bookings[index+1:]...)
*/
