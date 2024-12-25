package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	targetNumber := rand.Intn(100) + 1
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I've picked a number between 1 and 100. Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if guess < targetNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > targetNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number: %d\n", targetNumber)
			break
		}
	}
}
