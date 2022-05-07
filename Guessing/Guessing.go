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

const MAX = 100

func main() {
	rand.Seed(time.Now().UnixNano()) // set seed to let 'random'
	// variable initializing
	secret := rand.Intn(MAX)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("This is a guessing game, take a number from 0 to 100 (exclusive)")
	// infinite loop till correct
	for {
		input, err := reader.ReadString('\r') // Notice delimitation is different in different system
		if err != nil {
			fmt.Println("Unexpect error exists, please try guess again")
		}
		input = strings.TrimPrefix(input, "\n")
		input = strings.TrimSuffix(input, "\r")
		userInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please try guess again")
			continue
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
