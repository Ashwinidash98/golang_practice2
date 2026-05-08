package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const TotalPizzas = 8

var pizzasMade, pizzasFailed, total int

//Producer is a type for structs that holds two channels: one for pizzas, with all the
//informations for a given pizza order, whether it was made
///successfully, and another to handle end of processing(when we quit the channel)

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder is a type for structs that describes a given pizza order, It has the order
// number, a message indicating what happendedto the order, and a boolean
// indicating if the orderwas successfully completed.
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

//Close is simply a method to close the channel when we aredoen with it(i.e.
//something is pushed to the quit channel

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

//makePizza attempts to make a pizza, we generate a random number(1-12),
//and put two cases where we cannot make a pizza in time, Otherwise
//we make the pizza without issue. To make things interesting, each pizza
//willtake a different length of time to produce(some pizzas are harder than others).

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= TotalPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order %d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza %d. it will take %d seconds\n", pizzaNumber, delay)
		//delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("***we ran out of ingredients for pizza %d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("***the cook lquit while making the pizza %d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("pizza %d is ready", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

//pizzeria ois a Goroutine that runs in the background and
//callsmakePizza to try to make one order each time it iterates through the for loop.
//It iteratesuntil it receives something on the quitchannel.
//The quit channel does not receive anything until the consumer
//send it(when the Number of orders is greater than or equal to the constant TotalPizzas).

func pizzaria(pizzaMaker *Producer) {
	//keep track of which pizza we are making
	var i = 0
	//The loop will continue to execute, trying to make pizza
	// until the quit channel receive something

	//try to make pizza
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			//we tried to make a pizza(we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChannel := <-pizzaMaker.quit:
				//close channels
				close(pizzaMaker.data)
				close(quitChannel)
				return
			}
		}
		//try to make a pizza
		//decision
	}

}

func main() {
	//Seed a random number generator
	rand.Seed(time.Now().UnixNano())

	//print out a messange
	c := color.New(color.FgCyan)
	c.Println("The pizzaria is open for Business")
	c.Println("------------------------------------------")

	//Create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producre in the background
	go pizzaria(pizzaJob)

	//create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= TotalPizzas {
			if i.success {
				color.Green((i.message))
				color.Green("Order %d is out for delive!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The custmer is really mad!!")
			}
		} else {
			color.Cyan("DONE MAKING PIZZAS...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("****Error closing channel..")
			}
		}
	}

	//printout the ending message
	color.Cyan("DONE for the Day.")
	color.Cyan("we made %d pizzas, but failed to make %d pizzas, with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 7:
		color.Red("It was an awful Day")
	case pizzasFailed >= 5:
		color.Red("It could be a better Day")
	case pizzasFailed >= 3:
		color.Yellow("It was an okay Day")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good Day")

	default:
		color.Green("It was a great day!!")
	}

}
