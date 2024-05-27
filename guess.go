package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	const WELCOME = "Welcome to Guess game 🖥️"
	const GOODBYE = "\nGood bye ✌️🙏"

	fmt.Println(WELCOME)
	fmt.Println("------------------------")
	fmt.Println()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := rnd.Intn(100)

	var guesses = 0

	for {
		guesses += 1
		fmt.Println("Please enter your number. Press 'q' to quit.")
		var choice string

		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println(err)
			guesses -= 1
			continue
		}

		if choice == "q" {
			fmt.Println(GOODBYE)
			return
		}

		n, err := strconv.Atoi(choice)

		if err != nil {
			fmt.Println("Invalid number entered", choice)
			guesses -= 1
			continue
		}

		switch {
		case n == random:
			fmt.Printf("You guessed the right number %s in %d guesses 😊\n", choice, guesses)
			fmt.Println("Do you want to continue the game? Press any key to continue, 'q' to quit.")

			var ans string
			_, err := fmt.Scanln(&ans)

			if err == nil && ans == "q" {
				fmt.Println(GOODBYE)
				return
			}

			random = rnd.Intn(100)
			guesses = 0
			continue
		case n < random:
			fmt.Println("Your guess is less than the number that I thought. Guess again... 🤔")
			continue
		case n > random:
			fmt.Println("Your guess is more than the number that I thought. Guess again... 🤔")
			continue
		}
	}
}
