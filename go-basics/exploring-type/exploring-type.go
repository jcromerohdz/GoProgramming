package main

import "fmt"

var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I  ASSIGN the value "Shaken, not stirred"
// Go is a STATIC programming language
// Not a DYNAMIC programming language

var z string = "Shaken, not stirred"
var a   string = `Bond said, "Shaken, not stirred"`

func main(){
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}