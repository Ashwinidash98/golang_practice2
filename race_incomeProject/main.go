package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variables for bankbalance
	var bankBalance int

	//Printout starting values
	// \u20B9 is the Unicode for the Rupee symbol (₹)
	fmt.Printf("Initial Bank Balance: \u20B9%d.00", bankBalance)
	fmt.Println()

	//define weekly income
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "GIFT", Amount: 50},
		{Source: "Part-time job", Amount: 100},
		{Source: "Investments", Amount: 200},
	}
	wg.Add(len(incomes))

	//loop through 52 weeks and printout howmuch is made; keep a running total
	for i, incomes := range incomes {
		go func(i int, income Income) {

			defer wg.Done()
			for week := 1; week <= 52; week++ {
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp

				fmt.Printf("On week %d ,Income is %d.00, from %s", week, income.Amount, income.Source)
				fmt.Println()
			}
		}(i, incomes)
	}

	wg.Wait()

	//print out final balance
	fmt.Printf("final bank balance : %d.00", bankBalance)
	fmt.Println()
}
