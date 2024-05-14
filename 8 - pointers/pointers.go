package main

import "fmt"

func main() {
	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)

	variable1++
	fmt.Println(variable1, variable2)

	var variable3 int
	var pointer *int
	fmt.Println(variable3, pointer)

	variable3 = 100
	pointer = &variable3 // Store the memory address from variable3
	fmt.Println(variable3, pointer)

	variable3 = 150
	fmt.Println(variable3, *pointer) // *pointer = restore the value from the memory address

	a := 10
	b := &a
	fmt.Println(a, *b)
	a++
	fmt.Println(a, *b)
}
