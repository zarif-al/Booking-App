package printBookings

import (
	"Booking-App/model"
	"fmt"
)

func PrintBookings(bookings []model.User) {
	for i, bookings := range bookings {
		fmt.Println("User number: ", i)
		fmt.Println("FirstName: ", bookings.FirstName)
		fmt.Println("LastName: ", bookings.LastName)
		fmt.Println("Email: ", bookings.Email)
		fmt.Println("Number of Tickets: ", bookings.NumberOfTickets)
		fmt.Println("Venue: ", bookings.Venue)
	}
}