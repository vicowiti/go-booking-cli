package validateData

import (
	"fmt"
	"strings"
)

func ValidateData (firstName string, lastName string, email string, tickets uint, remainingTickets uint) bool{

	var isValidNames = len(firstName) > 1 && len(lastName) > 1
	var isValidEmail = strings.Contains(email,"@")
	
	var isValidTickets = tickets > 0 && tickets <= remainingTickets 
fmt.Print(isValidTickets)
	return isValidEmail && isValidNames && isValidTickets

}