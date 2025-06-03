package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate random between 1 and 100
	target := random.Intn(100) + 1

	fmt.Println("-- Guessing Game --")
	fmt.Println("I have chosen a number 1-100")
	fmt.Println("Can you guess it?")

	var guess int
	for {
		fmt.Println("Enter your guess")
		fmt.Scanln(&guess)

		// Check if guess is correct
		if guess == target {
			fmt.Println("You guessed the correctly!")
			break
		} else {
			if guess < target {
				fmt.Println("Try higher")
			} else {
				fmt.Println("Try lower")
			}
		}
	}
}
