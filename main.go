package main

//ctrl+c to stop application in commmandline terminal
import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Shubham Sharma"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	//arrays
	//var bookings = [50]string{"Shubham","amit","Sanjay","Arnav"}

	//var bookings [50]string

	//bookings[0]="Shubham"

	//slices

	var bookings []string

	//alternate declairation syntax for variables
	childName := "Amit"

	fmt.Printf(" conferenceTickets is %T, remainingTickets is %T \n", conferenceTickets, remainingTickets)
	fmt.Println(childName)
	fmt.Print("Hello ")
	fmt.Println(conferenceName)
	fmt.Println("You can book tickets for  ", conferenceName, " conference")
	fmt.Printf("Have you booked %v tickets?\n", conferenceName)
	fmt.Println("Total Tickets= ", conferenceTickets, " RemainingTickets= ", remainingTickets)

	for {
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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets

			//initialising arrays
			//bookings[0] = userName

			//initialising slices
			bookings = append(bookings, userName)

			//printing

			//fmt.Printf("This is first elm of booking arrays %v \n", bookings[0])

			//fmt.Printf("Array Type %T \n ", bookings)

			//fmt.Printf("Array length is %v \n", len(bookings))

			fmt.Printf("Thank you %v for booking %v ticket. You will recieve confirmation email on %v \n", userName, conferenceName, email)

			fmt.Printf("Tickets Remaining %v \n", remainingTickets)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("FirstName of booking are  %v \n", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Printf("Our Conference is booked OUT.")
				break
			}

		} else {
			fmt.Printf("We only have %v tickets remaining,so you can't book %v tickets \n", remainingTickets, userTickets)
			//continue
		}

	}

}
