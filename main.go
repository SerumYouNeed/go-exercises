package main

import (
	"fmt"
)

func main () {
	fmt.Println("Hello World!")

	var x int = 5
	fmt.Printf("The value of x is %d\n", x)

	for i := 0; i < 10; i++ {
		fmt.Printf("The value of i is %d\n", i)
	}	
}