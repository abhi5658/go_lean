package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("message is: ", quote.Go())
	fmt.Println(Hello("abhishek"))
}

func Hello(name string) string {
	message := fmt.Sprintf("hello %v. How are you?", name)
	return message
}
