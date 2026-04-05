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
	fmt.Println("Let's learn waitgroup!")

	var wg sync.WaitGroup

	words := []string{
		"Hii",
		"Hello",
		"Nmaste",
		"dhv",
		"sjjkl",
		"sdafsdas",
		"ADSDVFB",
		"SADFF",
	}
	wg.Add(len(words))

	for a, b := range words {
		go PrintSomething(fmt.Sprintf("%d: %s", a, b), &wg)
	}
	wg.Wait()
	wg.Add(1)

	PrintSomething("Let's see what is the result", &wg)
}
