package main

import (
	"fmt"
	"strings"
)

// It simply gets from the ping channel, converts it to uppercase,
// and appends a few exclamation marks and then sends the transformed text to the pong channel.
func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER(enter q to quit)")

	for {
		//print a promt
		fmt.Print(("-> "))

		//get user input
		var userInput string

		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break

		}
		ping <- userInput

		//wait for a responce
		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All Done. Closing channels.")
	close(ping)
	close(pong)
}
