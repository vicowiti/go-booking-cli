package captureData

import "fmt"

func CaptureData() (string,string,string,uint) {
	var firstName string
	var lastName string
	var email string
	var tickets uint

	fmt.Println("What is your first Name?")
	fmt.Scan(&firstName)

	fmt.Println("What is your last Name?")
	fmt.Scan(&lastName)

	fmt.Println("What is your email?")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&tickets)

	return firstName , lastName , email, tickets


}