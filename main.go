package main

import (
	"fmt"
)

func main() {

	var conferenceName string = "Akshay's Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// var bookings [50]string --> Array
	var bookings []string // --> Slice

	for len(bookings) <= 50 && remainingTickets > 0 {
		// ask user for their name
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidNumberOfUserTickets := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidNumberOfUserTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The first names of nookings are: %v\n", getFirstNames(bookings))

			if remainingTickets == 0 {
				fmt.Println("The concert is booked out. Please come back next year")
				break
			}

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, conferenceName)

		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address entered does not contain @")
			}
			if !isValidNumberOfUserTickets {
				fmt.Println("The number of tickets you want are more than available tickets")
			}

		}

	}

}
