package main

import (
	"fmt"
	"sort"
)

func main() {
	fruitList := []string{}

	fmt.Println(fruitList)
	fmt.Printf("Type of variable is %T\n", fruitList)

	// append an item
	fruitList = append(fruitList, "mango", "apple", "raspberry", "blueberry")
	fmt.Println("Length of Array: ", len(fruitList))
	fmt.Println("Spliced Array between 1-3:", fruitList[1:3])

	scores := make([]int, 3)
	scores[0] = 99
	scores[1] = 92
	scores[2] = 66

	fmt.Println("Scores before adding:", scores)

	scores = append(scores, 77, 88, 56, 99)

	fmt.Println("Scores After adding:", scores)
	fmt.Println("Is array sorterd:", sort.IntsAreSorted(scores))

	sort.Ints(scores)

	fmt.Println("Length of Array:", scores)
	fmt.Println("Is array sorted:", sort.IntsAreSorted(scores))

	// Remove an element in an array
	index := 1
	scores = append(scores[:index], scores[index+1:]...)
	fmt.Println("After removing First index:", scores)

}
