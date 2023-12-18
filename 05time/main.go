package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time is", currentTime)
	fmt.Println("Formatted Current time is", currentTime.Format("02/01/2006, Monday"))
	fmt.Println()

	customDate := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Custom date is", customDate)
	fmt.Println("Formatted Custom date is", customDate.Format("02/01/2006, Monday"))
}
