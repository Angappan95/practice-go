package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcomeMsg := "Welcome User"
	fmt.Println(welcomeMsg, "Enter rating for the Product")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Print("Thanks for rating ", input)
	fmt.Printf("Rating type is %T", input)

	parsedInput, err := strconv.ParseFloat(strings.TrimSpace(input), 32)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adjusted Rating is", parsedInput*2)
	}

}
