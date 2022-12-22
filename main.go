package main // Starting point of a go application.

import (
	"fmt"     //Importing a package from th go standard library
	"strings" //Provides methods for string manipulation
)

func main (){
var eventName = "Manchester Derby" // Variable declaration implicit type
const eventTickets uint = 86 //Constant, a value that does not change
var remainingTickets uint = 86 // Explicitly declaring a type 

fmt.Printf("Welcome to the %s\n", eventName) //Using printf to print formatted output
fmt.Println("Remaining tickets:", remainingTickets)
fmt.Println("Book your tickets here!")
fmt.Println("Seats :",eventTickets)
 bookings  := []string{} //Slice => dynamic array. Size not explicitly defined




for {//a for loop with just a body runs infinitely
var firstName string
var lastName string
var email string
var userTickets uint
//  Getting user input

// var bookings = [eventTickets]string{} //Declaring a string array


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

 bookings = append(bookings, firstName + " " + lastName) //Add item to a slice
fmt.Printf("%v remaining tickets\n", remainingTickets)

fmt.Printf("length of booking slice is: %v", len(bookings))//Length of  a slice
firstNameList := []string {}

for _, booking := range bookings { //A for-each loop // The underscore character is used as a blank identifier
	// Split  a sting by its white spaces
	var names = strings.Fields(booking)
	firstNameList = append(firstNameList, names[0])
}

fmt.Printf("The firstnames are: %v\n", firstNameList )

}

}