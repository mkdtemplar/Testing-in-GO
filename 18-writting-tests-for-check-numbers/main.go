package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	intro() // prints a welcome message

	// create a chanel to indacate when user want to quit
	doneChan := make(chan bool)

	// start the go routine
	go readUserInput(doneChan)

	// block until user quits
	<-doneChan

	// close the chanel
	close(doneChan)

	fmt.Println("Program exited")
}

func intro() {
	fmt.Println("Check if number is prime")
	fmt.Println("------------------------")
	fmt.Println("Enter integer to check if it is prime or q to quit")
	prompt()
}

func prompt() {
	fmt.Print("--> ")
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

	// check if user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert user input
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Enter integer!", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime by defenition", n)
	}
	if n < 0 {
		return false, "Negative numbers are not prime by definition"
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime number", n)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)
}
