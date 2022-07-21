package main

import "fmt"

func main() {

	conference_name := "Go Conference"
	const conference_tickets = 50
	var remaining_tickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conference_name)
	fmt.Println("Welcome to", conference_name, "booking application")

	fmt.Printf("Get your tickets here to attend. We have total %v tickets and %v tickets are left\n", conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend. We have total", conference_tickets, "tickets and", remaining_tickets, "tickets are left")

	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Ask the user for their name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole array : %v\n", bookings)
	fmt.Printf("The first value : %v\n", bookings[0])
	fmt.Printf("Slice type : %T\n", bookings)
	fmt.Printf("Slice length : %v\n", len(bookings))

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book")
	fmt.Scan(&userTickets)

	remaining_tickets -= userTickets

	fmt.Printf("%v %v with email %v has booked %v tickets\n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remaining_tickets, conference_name)
}
