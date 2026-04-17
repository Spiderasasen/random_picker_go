package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//getting the seed
	rand.Seed(time.Now().UnixNano())

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

	//gets a random option and returns it and prints it
	var selction string = random_option(options)
	fmt.Println("You got: ", selction)
}

// getting a random word from the list
func random_option(list []string) string {
	var index int = rand.Intn(len(list))
	return list[index]
}
