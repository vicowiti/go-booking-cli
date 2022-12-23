package main

import (
	"fmt"
	"go-booking-cli/captureData"
	"go-booking-cli/greeting"
	"go-booking-cli/validateData"
	"strconv"
)

var totalTickets uint = 80
var remainingTickets uint = 80
var eventName string = "Go Conference"
var bookings = make([]map[string] string,0)


func main () {
	// var firstName string
	// var lastName string
	// var email string
	// var tickets uint

	for {
		greeting.GreetUser(totalTickets, eventName)
	// greeting.Confirmation()


	var firstName , lastName , email, tickets = captureData.CaptureData()

	var validData = validateData.ValidateData(firstName, lastName,email, tickets, remainingTickets)

	if validData{
		remainingTickets = remainingTickets - tickets

		// Create a user Map

		var user = make(map[string] string) // Create a map for storing user data
		user["firstName"] = firstName
		user["lastName"] = lastName
		user["email"] = email
		user["tickets"] = strconv.FormatUint(uint64(tickets), 10) //Converting uint to string

		bookings = append(bookings, user)
		var firstNameSlice = []string{}

		for _, booking := range bookings {
			firstNameSlice = append(firstNameSlice, booking["firstName"])
		}

		fmt.Printf("The user list is: %v \n", bookings)
		fmt.Printf("The user list is: %v \n", firstNameSlice)
		fmt.Printf("Total remaining tickets is: %v\n", remainingTickets)
	}else{
		fmt.Println("You entered invalid data")
		fmt.Println("Kindly confirm the following:")
		fmt.Println("1. Your names are at least two characters long.")
		fmt.Println("2. Your email contains the '@' symbol.")
		fmt.Println("3. You are booking a valid number of tickets.")
	}
	}
}