package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("hello, %v. Welcome", name)
	return message
}

func ValidateAndHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("hello, %v. Welcome", name)
	return message, nil
}
