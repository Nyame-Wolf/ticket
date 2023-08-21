package helper

import "strings"

func ValidateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	//user validation: names atleast 2 char long
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	//validate email
	isValidEmail := strings.Contains(email, "@")
	//validate tickets
	isValidTicketNo := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNo
}
