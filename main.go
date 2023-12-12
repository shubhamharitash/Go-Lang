package main

import "fmt"

func main() {

	var conferenceName = "Shubham Sharma"
	const conferenceTickets = 50
	var remainingTickets = 50

	//alternate declairation syntax for variables
	childName := "Amit"

	fmt.Printf(" conferenceTickets is %T, remainingTickets is %T \n", conferenceTickets, remainingTickets)
	fmt.Println(childName)
	fmt.Print("Hello ")
	fmt.Println(conferenceName)
	fmt.Println("You can book tickets for  ", conferenceName, " conference")
	fmt.Printf("Have you booked %v tickets?\n", conferenceName)
	fmt.Println("Total Tickets= ", conferenceTickets, " RemainingTickets= ", remainingTickets)

	var userName string
	var userTickets int
	var email string

	//ask user for conferenceName

	fmt.Println("Enter UserName")
	fmt.Scan(&userName)

	fmt.Println("Enter Email Id")
	fmt.Scan(&email)

	fmt.Println("Enter No of Tickets to Book")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v for booking %v ticket. You will recieve confirmation email on %v \n", userName, conferenceName, email)

	fmt.Printf("Tickets Remaining %v \n", remainingTickets)

}
