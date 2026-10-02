package main

import "fmt"

func main(){
	// ; is added authomatically at the end
	var channelName string
	channelName = "Learn Golang"
	var year int = 2026

	fmt.Println("Chanel:", channelName)
	fmt.Println("Year:", year)

	var rating float64 = 3.4
	fmt.Println("Rating:", rating)

	var isActive = true // type is isferred from the value
	fmt.Println("Is Active:", isActive)

	guestCount := 10 // short declaration, type is inferred
	fmt.Println("Guest Count:", guestCount)

	if guestCount == 10 {
		fmt.Println("Yep!")
	}

	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}

	text := "Hello"
	fmt.Printf("Type of character: %T\n", text[1]) // uint8
	fmt.Printf("Type of character: %T\n", []rune(text)[1]) // rune -> int32
}