package main

import "fmt"

var Country = "India"

func main() {
	var username string = "Angappan"
	var isStudent bool = false
	var age int8 = 28
	fmt.Println("Username is", username)
	fmt.Println("Status is", isStudent)
	fmt.Println("Age is", age)

	var city = "Chennai"
	fmt.Println("Location is", city)

	score := 95.2
	fmt.Println("Score is", score)
	fmt.Printf("Score is of type %T", score)

	fmt.Println()
	fmt.Println("Country is", Country)
}
