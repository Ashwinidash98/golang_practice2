package main

import (
	"fmt"
	"sync"
	"time"
)

//The Dining Philosophers problem is well known in comouter science circles.
//5 philosophores, numbered from 0 to 4, live in a house where the
//table is laid for tehm; each philosopher has their own place at the table.
//Their only difficulty -besides those of philosopher- is that the dish
//served is a very differeent kind of spagetti which has to be eaten with
//two forks. There are two forks next to each pate, so that presents no
//dificulty. As a concequence, however, this means that no two neighbours
//may e eating simultaneously, since there are five philosophoers and five forks.
//
//This is a simple implementation of Dijkstra's solution to the "Dining
// Philosopher's Dilemma"

// Philosopher is a struct which stres info about a philosopher
type Philosopher struct {
	name     string
	rihtFrk  int
	leftFork int
}

// philosophores is a list of all
var philosophoers = []Philosopher{
	{name: "Plato", rihtFrk: 0, leftFork: 4},
	{name: "Socrates", rihtFrk: 1, leftFork: 0},
	{name: "Aristotle", rihtFrk: 2, leftFork: 1},
	{name: "Pascal", rihtFrk: 3, leftFork: 2},
	{name: "Locke", rihtFrk: 4, leftFork: 3},
}

// define some variables
var hunger = 3 //how many times does a person eat
var eatTime = 1 * time.Second
var think = 3 * time.Second
var sleep = 1 * time.Second

func main() {
	//print out a welsome msg
	fmt.Println("Dining Philosphers Problem")
	fmt.Println("------------------------------------")
	fmt.Println("The table is empty.")

	//start meal
	dine()

	// printout finished msg
	fmt.Println("The table is empty.")

}
func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophoers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophoers))

	//forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophoers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//start the meal
	for i := 0; i < len(philosophoers); i++ {
		//fire off a goroutine for the current philosopher
		go diningProblem(philosophoers[i], wg, forks, seated)
	}
	wg.Wait()
}

// diningProblem is the function fired off aas a goroutine for each of our philosphers. It takes one
// philosopher, our Waitgroup to determine when everyone is done, a map containing the mutexes for every
// fork on the table, and a waitgroup used to pause execution of ebvery instance of this goroutine
// until everyone is seated at the table.
func diningProblem(Philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	//seat the philosopher at the table
}
