package main

import (
	"fmt"

	"rsc.io/quote"
)

func Hello(name string) string { // exported as function name starts with capital letter
	message := fmt.Sprintf("hello %v. How are you?", name)
	return message
}

func add(x, y int) int { // unexported as function name starts with small letter
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
func nakedFunc(sum int) (x, y int) {
	x = sum / 4
	y = sum + 4
	return
}

func main() {
	fmt.Println("message is: ", quote.Go())
	fmt.Println(Hello("abhishek"))
	fmt.Println("sum for 2 and 3 is: ", add(2, 3))
	a, b := swap("string-a", "string-b")
	fmt.Println("after swapped", a, b)
	c, d := nakedFunc(16)
	fmt.Println("naked return is:", c, d)
}

/*
$> go run .
message is:  Don't communicate by sharing memory, share memory by communicating.
hello abhishek. How are you?
sum for 2 and 3 is:  5
after swapped string-b string-a
naked return is: 4 20
*/
