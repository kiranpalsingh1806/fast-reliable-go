package main

import "fmt"

func main() {

	var conference_name = "Go Conference"
	const conference_tickets = 50
	var remaining_tickets = 50

	fmt.Printf("Welcome to %v booking application\n", conference_name)
	fmt.Println("Welcome to", conference_name, "booking application")

	fmt.Printf("Get your tickets here to attend. We have total %v tickets and %v tickets are left\n", conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend. We have total", conference_tickets, "tickets and", remaining_tickets, "tickets are left")

	var user_name string
	var user_tickets int
	// Ask the user for their name
	user_name = "Tom"
	user_tickets = 2

	fmt.Printf("User %v has booked %v tickets\n", user_name, user_tickets)
}
