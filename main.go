package main

import (
	"fmt"
)

func main() {
	//makes an emptry array
	var options []string

	//user makes the choices
	var choice string
	for {
		fmt.Println("Type stop to end.")
		fmt.Println("Please enter a option: ")
		fmt.Scan(&choice)

		//if the user types stops, it will end the loop
		if choice == "stop" || choice == "Stop" {
			break
		}

		//otherwise it will appended to the list
		options = append(options, choice)
	}

}
