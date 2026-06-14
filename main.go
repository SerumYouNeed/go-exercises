package main

import (
	"fmt"
)

func main () {
	fmt.Println("Hello World!")

	m := [3]string{"Go", "is", "awesome"}

	for _, word := range m {
		fmt.Printf("%s ", word)
	}

}