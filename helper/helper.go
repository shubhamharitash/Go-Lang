package helper

import "strings"

//capitalize func name to export it outside package
func ValidateUserInput(userName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidName, isValidTicketNumber
}
