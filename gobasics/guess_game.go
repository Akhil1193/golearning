package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func game() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1
	//welcome message
	fmt.Println("Welcome to guessing game!")
	fmt.Println("I have guessed a number between 1 to 100")
	fmt.Println("It's time for you to choose!")

	var guess int
	for {
		fmt.Println("Enter your guess:")
		fmt.Scan(&guess)
		if guess < target{
			fmt.Println("You have guessed a lower number! Please try again!")
		} else if guess > target{
			fmt.Println("You have guessed a higher number! Please try again!")
		} else if guess == target{
			fmt.Println("Congratulations! You have guessedd the number correctly!")
			break
		}
	}

}
