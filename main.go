package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets int = 25
	var remainingTickets uint = 25

	var bookings []string
	conferenceName := "Virtual Conference"

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("In the next step you can get your tickets here to attend.")

	// Listen to user input
	for {
		var userFirstName string
		var userLastName string
		var userEmail string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&userFirstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&userLastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&userEmail)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userFirstName+" "+userLastName)

		fmt.Printf("Thank you %v %v for booking %v tickets! You will receive a confirmation email at %v soon.\n", userFirstName, userLastName, userTickets, userEmail)
		fmt.Printf("There are still %v tickets remaining for the %v.\n", remainingTickets, conferenceName)

		// Foreach loop (for + range)
		var firstNames []string
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is sadly booked out. Join us next year.")
			break
		}
	}
}
