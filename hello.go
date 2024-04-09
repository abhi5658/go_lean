package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
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

var abd = 123 // works
// xyz := 123 // does not work
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

func main() {
	fmt.Println("message is: ", quote.Go())
	fmt.Println("random number is: ", rand.Intn(100000))
	fmt.Println(Hello("abhishek"))
	fmt.Println("sum for 2 and 3 is: ", add(2, 3))
	a, b := swap("string-a", "string-b")
	fmt.Println("after swapped", a, b)
	c, d := nakedFunc(16)
	fmt.Println("naked return is:", c, d)

	var i int
	var e, java, python bool
	fmt.Println("declared variables", i, e, java, python)
	var lang = "go-lang" // no type needed, will deduce from initialised value
	var f, fifa = false, true
	var g, h int = 1, 2
	fmt.Println("initialised vars", lang, f, fifa, g, h)
	abc := 234 // := short assignment statement can be used in place of a var declaration with implicit type.
	fmt.Println("short assignment decalred variable abc:", abc)
	var uninitialisedInteger int
	var uninitialisedFloat float64
	var uninitialisedBoolean bool
	var uninitialisedString string
	fmt.Printf("uninitialised variables: %v %v %v %q\n", uninitialisedInteger, uninitialisedFloat, uninitialisedBoolean, uninitialisedString)

	var (
		ToBe     bool       = false
		MaxInt   uint64     = 1<<64 - 1
		zComplex complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Typre: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", zComplex, zComplex)

	var x, y int = 3, 4
	//error: cannot use (x * x + y * y) (value of type int) as float64 value in argument to math.Sqrt
	//var squareRoot float64 = math.Sqrt((x*x + y*y))
	var squareRoot float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(squareRoot)
	fmt.Println("conversion", x, y, z)
}

/*
$> go run .
message is:  Don't communicate by sharing memory, share memory by communicating.
random number is:  628
hello abhishek. How are you?
sum for 2 and 3 is:  5
after swapped string-b string-a
naked return is: 4 20
declared variables 0 false false false
initialised vars go-lang false true 1 2
short assignment decalred variable abc: 234
uninitialised variables: 0 0 false ""
Type: bool Value: false
Typre: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
conversion 3 4 5

*/
