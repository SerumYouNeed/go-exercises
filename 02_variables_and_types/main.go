package main

import "fmt"

func main(){
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
}