package main

import (
	"fmt"
	"sync"
)

func PrintSomething(s string, wg *sync.WaitGroup) {
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
		PrintSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)
	PrintSomething("Here is the waitgroup working", &wg)
}
