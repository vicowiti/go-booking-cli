package main

import "go-booking-cli/greeting"

var totalTickets uint = 80
var remainingTickets uint = 80
var eventName string = "Go Conference"

func main () {
	
	greeting.GreetUser(totalTickets, eventName)
	greeting.Confirmation()

}