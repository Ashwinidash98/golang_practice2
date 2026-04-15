package main

import (
	"fmt"
	"strings"
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
	name      string
	rightFork int
	leftFork  int
}

// philosophores is a list of all
var philosophoers = []Philosopher{
	{name: "Plato", rightFork: 0, leftFork: 4},
	{name: "Socrates", rightFork: 1, leftFork: 0},
	{name: "Aristotle", rightFork: 2, leftFork: 1},
	{name: "Pascal", rightFork: 3, leftFork: 2},
	{name: "Locke", rightFork: 4, leftFork: 3},
}

// define some variables
var hunger = 3 //how many times does a person eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

// ***added this
var orderMutex sync.Mutex  //a mutex for the slice orderFinished;
var orderFinished []string // the order in which philosophers finish dining and leave

func main() {
	//print out a welsome msg
	fmt.Println("Dining Philosphers Problem")
	fmt.Println("------------------------------------")
	fmt.Println("The table is empty.")

	//***added this
	time.Sleep(sleepTime)

	//start meal
	dine()

	// printout finished msg
	fmt.Println("The table is empty.")

	//***added this
	time.Sleep(sleepTime)
	fmt.Print("\nOrder finished:  ", strings.Join(orderFinished, " , "))

}
func dine() {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	//wait group
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophoers))

	//We want everyone to be seated before they start eating, so create a WaitGroup for that and set it to 5.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophoers))

	//forks is a map of all 5 forks. Forks are assigned using the fields leftFork and rightFork in the Philosopher type.
	//Each fork, then, can be found using the index (an integer), and each forlk has unique mutex
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophoers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//start the meal by iterating through our slice of Philosophers.
	for i := 0; i < len(philosophoers); i++ {
		//fire off a goroutine for the current philosopher
		go diningProblem(philosophoers[i], wg, forks, seated)
	}

	//Wait for the Philosopher to finish. This blocks until the wait group is 0
	wg.Wait()
}

// diningProblem is the function fired off aas a goroutine for each of our philosphers. It takes one
// philosopher, our Waitgroup to determine when everyone is done, a map containing the mutexes for every
// fork on the table, and a waitgroup used to pause execution of ebvery instance of this goroutine
// until everyone is seated at the table.
func diningProblem(Philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	//seat the philosopher at the table
	fmt.Printf("%s is seated at the table", Philosopher.name)

	//Decrement the seated Waitgroup by one
	seated.Done()

	//Wait until everyone is seated
	seated.Wait()

	//eat, think, hunger three times
	for i := hunger; i > 0; i-- {
		//get a lock on both sides
		if Philosopher.leftFork > Philosopher.rightFork {
			forks[Philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", Philosopher.name)
			forks[Philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", Philosopher.name)
		} else {
			forks[Philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", Philosopher.name)
			forks[Philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", Philosopher.name)
		}

		// forks[Philosopher.leftFork].Lock()
		// fmt.Printf("\t%s takes the left fork.\n", Philosopher.name)
		// forks[Philosopher.rightFork].Lock()
		// fmt.Printf("\t%s takes the right fork.\n", Philosopher.name)

		//by this time we get to this, the philosophore has a lock(Mutex) on both sides
		fmt.Printf("\t%s has both the forks\n", Philosopher.name)
		time.Sleep(eatTime)

		//The philosophore starts to think, but does not drop the forks yet.
		fmt.Printf("\t%s is thinking\n", Philosopher.name)
		time.Sleep(thinkTime)

		//Unlock the mutexes for both forks
		forks[Philosopher.leftFork].Unlock()
		forks[Philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", Philosopher.name)

	}
	//The philosophore has finished eating, so print out a message.
	fmt.Println(Philosopher.name, "--------------------------")
	fmt.Println(Philosopher.name, "is satisfied.")
	fmt.Println(Philosopher.name, "left the table.")
	fmt.Println(Philosopher.name, "--------------------------")

	//***added this
	orderMutex.Lock()
	orderFinished = append(orderFinished, Philosopher.name)
	orderMutex.Unlock()
}
