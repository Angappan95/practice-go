package main

import "fmt"

func main() {
	var fruitList [3]string

	fruitList[0] = "kivi"
	fruitList[1] = "orange"

	fmt.Println(fruitList)
	fmt.Printf("Type of fruitList %T\n", fruitList)

	vegList := [3]string{"potato", "onion"}
	fmt.Println(vegList)

}
