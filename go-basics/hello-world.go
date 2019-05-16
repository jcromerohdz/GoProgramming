package main

import "fmt"

func main()  {
	fmt.Println("Hello Christian, How is you day today")
	foo()
	fmt.Println("Something more")

	for i := 0; i<100; i++{
		if i %2 == 0{
		  fmt.Println(i)		
		}	
	}

	bar()
}

func foo(){
	fmt.Println("I'm in foo")
}

func bar(){
	fmt.Println("and the we exited!")
}

// control flow:
// (1) sequence
// (2) loop; interative
// (3) conditional