package main

import (
	"fmt"
	"strconv"
	"ticketing/helper"
)

// create package level variables
const conferenceTickets = 50
const conferenceName = "Go Conference"

// create an empty slice of maps bookings := []string{}
// add a second arg to make, for initial size of the map
var bookings = make([]map[string]string, 0)

// create  a variable to track available tickets
var remainingTickets uint = conferenceTickets

func main() {

	greetUsers()

	// indefinite loop to keep prompting user to book tickets
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNo := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// break from loop if user books tickets than what is available
		if isValidName && isValidEmail && isValidTicketNo {

			bookTickets(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("This are all of our bookings %v\n", firstNames)

			// use break keyword to break out of the loop
			/*can use this sytax in if
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining{
				...
			}
			*/
			if remainingTickets == 0 {
				fmt.Println("The conference is booked out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The name you entered is invalid. Both the first and last names must be atleast two characters long")
			}
			if !isValidEmail {
				fmt.Println("You entered an invalid email")
			}
			if !isValidTicketNo {
				fmt.Printf("We only have %v tickets available. Please book within range\n", remainingTickets)
			}
		}

	}

}

func greetUsers() {
	// create  welcome messages for conference attendees.
	fmt.Printf("Welcome to the %v.\nThere are a total of  %v tickets and only %v tickets remaining\n", conferenceName, remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here\n")
}

/*
// prints firstnames
	func printFirstNames(bookings []string) {
		// return slice with only firstNames:
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("This are all of our bookings %v\n", firstNames)
	}
*/

// returns a slice of firstnames. Print happens in main func.
func getFirstNames() []string {
	// return slice with only firstNames:
	// since its a map, remove string.Fields
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// get back user unput
	fmt.Println("Enter your first name to book:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName, lastName, email string) {
	// calculate remaining tickets:
	remainingTickets = remainingTickets - userTickets

	// create a map
	userData := make(map[string]string)
	//add data to the map
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//store new users in slice
	bookings = append(bookings, userData)

	fmt.Printf("Thank you  %v %v for booking %v tickets. You will receive a confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Remaining tickets\n\n", remainingTickets)

	// print list of bookings
	fmt.Printf("Here is a list of your bookings: %v\n", bookings)

}
