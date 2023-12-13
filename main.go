package main

//ctrl+c to stop application in commmandline terminal
import (
	"Booking-app/helper"
	"fmt"
	"strings"
)

// package-level variables
const conferenceTickets = 50

var conferenceName = "Go-Lang"
var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		userName, userTickets, email := getUserInput()

		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUserInput(userName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, userName, email)

			//call printFirstNames func
			firstNames := getFirstNames()
			fmt.Printf("FirstName of booking are  %v \n", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Printf("Our Conference is booked OUT.")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("UserName entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email doesn't contains @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid TicketNumber")
			}

			//fmt.Printf("Invalid Input, Enter input Correctly\n")
			//continue
		}

	}

}

func greetUsers() {
	fmt.Printf("welcome to %v Booking Application\n", conferenceName)
	fmt.Println("Total Tickets= ", conferenceTickets, " RemainingTickets= ", remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, uint, string) {
	var userName string
	var userTickets uint
	var email string

	//ask user for conferenceName

	fmt.Println("Enter UserName")
	fmt.Scan(&userName)

	fmt.Println("Enter Email Id")
	fmt.Scan(&email)

	fmt.Println("Enter No of Tickets to Book")
	fmt.Scan(&userTickets)

	return userName, userTickets, email
}

func bookTicket(userTickets uint, userName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userName)
	fmt.Printf("Thank you %v for booking %v ticket. You will recieve confirmation email on %v \n", userName, conferenceName, email)
	fmt.Printf("Tickets Remaining %v \n", remainingTickets)

}

//arrays
//var bookings = [50]string{"Shubham","amit","Sanjay","Arnav"}
//var bookings [50]string
//bookings[0]="Shubham"
//initialising arrays
//bookings[0] = userName

//slices
//alternate declairation syntax for variables
//childName := "Amit"

//print syntax
//fmt.Printf(" conferenceTickets is %T, remainingTickets is %T \n", conferenceTickets, remainingTickets)
//fmt.Print("Hello ")
//fmt.Println("You can book tickets for  ", conferenceName, " conference")
//fmt.Printf("Have you booked %v tickets?\n", conferenceName)
//fmt.Printf("This is first elm of booking arrays %v \n", bookings[0])
//fmt.Printf("Array Type %T \n ", bookings)
//fmt.Printf("Array length is %v \n", len(bookings))
