package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "ABC"
	expectedOutput := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := ValidateAndHello(name)
	if !expectedOutput.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(name) = %q, %v, want match for %#q, nil`, msg, err, expectedOutput)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := ValidateAndHello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
