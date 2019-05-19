package main

import "fmt"

// DECALRE the variable "y"
// ASSIGN the value 43
// declare & assign = initilization
var y = 43

// DECLARE ther is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSINGS the ZERO VALUE of TYPE int to "z"
// false for booleans, o for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main(){
	// short declaration operator
	// DECLARE a variabe and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)
	
	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo(){
	fmt.Println("again: ", y)
}

