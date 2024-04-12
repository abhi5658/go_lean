package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("hi_module:")
	log.SetFlags(0)

	//message := greetings.Hello("")
	//fmt.Println(message)
	message2, err := greetings.ValidateAndHello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message2)
}
