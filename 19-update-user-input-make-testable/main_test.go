package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{
			name:     "prime",
			testNum:  7,
			expected: true,
			msg:      "7 is a prime number",
		},
		{
			name:     "not prime",
			testNum:  8,
			expected: false,
			msg:      "8 is not prime number",
		},
		{
			name:     "zero",
			testNum:  0,
			expected: false,
			msg:      "0 is not prime by definition",
		},
		{
			name:     "one",
			testNum:  1,
			expected: false,
			msg:      "1 is not prime by definition",
		},
		{
			name:     "negative",
			testNum:  -4,
			expected: false,
			msg:      "Negative numbers are not prime by definition",
		},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s, but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy od os.Stdout
	old := os.Stdout

	// create read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close the writer
	_ = w.Close()

	// reset os.Stdout() to previous
	os.Stdout = old

	// read the output from prompt() function
	out, _ := io.ReadAll(r)

	if string(out) != "--> " {
		t.Errorf("expected --> , but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy od os.Stdout
	old := os.Stdout

	// create read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close the writer
	_ = w.Close()

	// reset os.Stdout() to previous
	os.Stdout = old

	// read the output from prompt() function
	out, _ := io.ReadAll(r)
	expected := "Check if number is prime\n------------------------\nEnter integer to check if it is prime or q to quit\n--> "
	got := string(out)

	if expected != got {
		t.Errorf("expected %s , but got %s", expected, got)
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Enter integer!"},
		{name: "zero", input: "0", expected: "0 is not prime by definition"},
		{name: "one", input: "1", expected: "1 is not prime by definition"},
		{name: "two", input: "2", expected: "2 is a prime number"},
		{name: "quit", input: "q", expected: ""},
		{name: "three", input: "3", expected: "3 is a prime number"},
		{name: "negative", input: "-4", expected: "Negative numbers are not prime by definition"},
		{name: "eight", input: "8", expected: "8 is not prime number"},
	}

	for _, tt := range tests {
		input := strings.NewReader(tt.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(res, tt.expected) {
			t.Errorf("%s: expected %s, but got %s", tt.name, tt.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this function we need a chaneel and instance of io.Reader
	doneChan := make(chan bool)

	// create reference to bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
}
