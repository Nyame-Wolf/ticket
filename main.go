package main

import (
	"fmt"
)

// create package level variables
const conferenceTickets = 50
const conferenceName = "Go Conference"

func main() {
	// create  a variable to track available tickets
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var remainingTickets uint = conferenceTickets

	//create empty slice bookings := []string{}
	var bookings []string

	// create a welcome message for conference attendees.
	fmt.Printf("Welcome to the %v. There are a total of  %v tickets and only %v tickets remaining\n", conferenceName, remainingTickets, conferenceTickets)

	// get back user unput
	fmt.Println("Enter your first name to book:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	// calculate remaining tickets:
	remainingTickets = remainingTickets - userTickets

	//store new users in slice
	bookings = append(bookings, firstName+" "+lastName)

	// see test bookings in slice
	fmt.Printf("Slice value %v\n", bookings)
	fmt.Printf("The type is %T\n", bookings)

	fmt.Printf("Thank you  %v %v for booking %v tickets. You will receive a confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Remaining tickets\n", remainingTickets)

	fmt.Printf("This are all of our bookings %v", bookings)

}
