package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	n := 0
	_, msg := isPrime(n)
	fmt.Println(msg)
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
