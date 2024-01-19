package main

import (
	"io"
	"os"
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
			msg:      "0 is not prime by defenition",
		},
		{
			name:     "one",
			testNum:  1,
			expected: false,
			msg:      "1 is not prime by defenition",
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
	//
}
