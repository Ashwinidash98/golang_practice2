package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"delta",
		"beta",
		"gamma",
		"gfdesefr",
		"fgrtyhg",
		"gamyttdma",
		"hytgfd",
		"gamkjhgma",
	}
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("Here is the waitgroup working", &wg)
}
