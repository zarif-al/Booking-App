package welcome

import (
	"fmt"
)

func WelcomeMsgPrint(conferenceName string)  {
	fmt.Printf("Welcome to %v conference ticket booking application.\n", conferenceName)
	fmt.Println("Please provide your necessary details to book your ticket.")
}