package main

import (
	"fmt"
)

func main () {

	for _, v := range []int{1, 2, 3, 4, 5, 6} {

		fmt.Printf("%d ", v)
		if v == 5 {
			continue
		}
	}

	z := 1
	if z > 0 {
		fmt.Println("z is bigger than 0")
}
}
