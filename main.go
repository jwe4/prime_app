package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// print a welcome message
	intro()

	// create a channel to indicate when the user want sot quit
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(doneChan)

	// block until the done channel gets a value
	<-doneChan
	// close the channel
	close(doneChan)
	// say goodbye
	fmt.Println("Goodbye.")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user typed into an int
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(num)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	} else if n < 0 {
		return false, fmt.Sprintf("Negative numbers are not prime, by definition!")
	}
	// use the modulus operator, repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime, it is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number!", n)
}
