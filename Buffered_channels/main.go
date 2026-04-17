package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		//print a got data msg
		i := <-ch
		fmt.Println("Got", i, "from channel")

		//simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//ch := make(chan int)//unbuffered
	ch := make(chan int, 5) //buffered

	go listenToChan(ch)

	for i := 0; i <= 7; i++ {
		fmt.Println("sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel !!")
	}

	fmt.Println("Done!!")
	close(ch)
}
