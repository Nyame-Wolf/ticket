package main

import (
	"fmt"
)

// create package level variables
const conferenceTickets = 50
const conferenceName = "Go Conference"

func main() {
	// create  a variable to track available tickets
	var userTickets uint
	var remainingTickets uint = conferenceTickets - userTickets

	// create a welcome message for conference attendees.
	fmt.Printf("Welcome to the %v. There are a total of  %v tickets and only %v tickets remaining\n", conferenceName, remainingTickets, conferenceTickets)
	fmt.Println("Pick your tickets here to attend")

}
