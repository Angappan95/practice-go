package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMsg := "Welcome User"
	fmt.Println(welcomeMsg, "Enter rating for the Product")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Print("Thanks for rating ", input)
	fmt.Printf("Rating type is %T", input)

}
