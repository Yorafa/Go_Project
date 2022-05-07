package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXV2 = 100

func main() {
	rand.Seed(time.Now().UnixNano()) // set seed to let 'random'
	// variable initializing
	secret := rand.Intn(MAXV2)
	var userInput int
	fmt.Println("This is a guessing game, take a number from 0 to 100 (exclusive)")
	// infinite loop till correct
	for {
		n, err := fmt.Scanf("%d\r\n", &userInput) // Notice delimitation is different in different system
		if err != nil || n == 0 || userInput >= 100 {
			fmt.Println("Invalid input, please try guess again")
		}
		if userInput == secret {
			fmt.Println("Correct guess, congratulation!!")
			break
		} else if userInput < secret {
			fmt.Println("your guess is less than the number")
		} else if userInput > secret {
			fmt.Println("your guess is larger than the number")
		}
	}
}
