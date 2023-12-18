package main

import "fmt"

func main() {
	fmt.Println("Practicing Pointers")

	var ptr *string

	var name = "Ram"

	ptr = &name

	fmt.Println("Name is", name)
	fmt.Println("Pointer value is", ptr)
	fmt.Println("Value of Pointer is", *ptr)

	*ptr = *ptr + " Kannappan"
	fmt.Println("Pointer address after concatenation", ptr)
	fmt.Println("Value of ptr after concatenation", *ptr)

	fmt.Println("Testing -> *&", *&name)

}
