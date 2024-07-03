package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	success := false

	for i := 10; i > 0; i-- {
		fmt.Println("You have", i, "guesses left")

		fmt.Println("Make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess == target {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		} else if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else {
			fmt.Println("Oops. Your guess was HIGH")
		}
	}

	if !success {
		fmt.Println("Sorry. You didn't guess my number. It was", target)
	}
}
