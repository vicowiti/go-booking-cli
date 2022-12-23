package greeting

import (
	"fmt"
)

func GreetUser (totalTickets uint, eventName string) {

	fmt.Printf("Welcome to our %v Ticket Booking Application!\n", eventName)
	fmt.Printf("We have a total of %v tickets.\n", totalTickets)
	// fmt.Println("Do you want to proceed to book tickets? Yes/No")

}


// func Confirmation () {
// 	var isAffirmative string

// 	fmt.Scan(&isAffirmative) //Get user input and assign it to string variable

// 	// Convert input to lowercase

// 	isAffirmative = strings.ToLower(isAffirmative)
// 	//Check for answer
// 	for isAffirmative != "yes" {
// 		fmt.Println("You chose to exit the application")
// 		break
// 	}
	


// }