package main

import "fmt"

func main (){
var eventName = "Manchester Derby" // Variable declaration
const eventTickets uint = 86 //Constant, a value that does not change
var remainingTickets uint = 86

fmt.Printf("Welcome to the %s\n", eventName) //Using printf to print formatted output
fmt.Println("Remaining tickets:", remainingTickets)
fmt.Println("Book your tickets here!")
fmt.Println("Seats :",eventTickets)

var firstName string
var lastName string
var email string
var userTickets uint
//  Getting user input

// var bookings = [eventTickets]string{} //Declaring a string array
var bookings = []string{} //Slice => dynamic array.

fmt.Println("What is your first name")
fmt.Scan(&firstName) //Scanning and assigning to the memory address

fmt.Println("What is your last name")
fmt.Scan(&lastName)

fmt.Println("What is your email")
fmt.Scan(&email) 

fmt.Println("How many tickets do you want?")
fmt.Scan(&userTickets)

fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v\n", firstName,lastName,userTickets,email)
 remainingTickets -=  userTickets

 bookings = append(bookings, firstName + " " + lastName)
fmt.Printf("%v remaining tickets\n", remainingTickets)

fmt.Printf("length of booking slice is: %v", len(bookings))
fmt.Printf("arr slice is: %v",bookings)
}