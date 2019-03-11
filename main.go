package main

import (
	"fmt"
	"math/rand"
)

// Attempt to reimplement the simulations that produced the numbers in
// Table 1 of "The Economic Limits of Bitcoin and the Blockchain" by
// Eric Budish
//
// https://faculty.chicagobooth.edu/eric.budish/research/Economic-Limits-Bitcoin-Blockchain.pdf

func main() {
	fmt.Println("e=0")
	fmt.Printf("A: %v; blocks: %v\n", 1.05, sim(1.05, 0))
	fmt.Printf("A: %v; blocks: %v\n", 1.1, sim(1.1, 0))
	fmt.Printf("A: %v; blocks: %v\n", 1.2, sim(1.2, 0))
	fmt.Printf("A: %v; blocks: %v\n", 1.25, sim(1.25, 0))
	fmt.Printf("A: %v; blocks: %v\n", 1.33, sim(1.33, 0))
	fmt.Printf("A: %v; blocks: %v\n", 1.5, sim(1.5, 0))
	fmt.Printf("A: %v; blocks: %v\n", 2, sim(2, 0))
	fmt.Printf("A: %v; blocks: %v\n", 5, sim(5, 0))
	fmt.Println("e=6")
	fmt.Printf("A: %v; blocks: %v\n", 1.05, sim(1.05, 6))
	fmt.Printf("A: %v; blocks: %v\n", 1.1, sim(1.1, 6))
	fmt.Printf("A: %v; blocks: %v\n", 1.2, sim(1.2, 6))
}

func sim(ax float64, escrow float64) float64 {
	n := 1000000
	var btotal float64
	for i := 0; i < n; i++ {
		btotal = btotal + one(ax, escrow)
	}
	return btotal / float64(n)
}

func one(ax float64, escrow float64) float64 {
	var honest_time float64
	var attacker_time float64
	var e float64
	for i := 0; i < 1000000; i++ {
		if e > escrow+1 && honest_time > attacker_time {
			return e
		}
		h := rand.ExpFloat64()
		honest_time = honest_time + h
		a := rand.ExpFloat64() / ax
		attacker_time = attacker_time + a
		e = e + 1
	}
	return 0
}
