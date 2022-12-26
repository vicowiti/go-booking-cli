package bookTicket

import (
	"fmt"
	"time"
)

func BookTicket (name string, tickets uint) {

	time.Sleep(3 * time.Second)

	fmt.Println("########################################################")
	fmt.Printf("Congratulations %v, You booked %v tickets\n", name, tickets)
	fmt.Println("########################################################")


}