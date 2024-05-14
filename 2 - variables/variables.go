package main

import "fmt"

func main() {
	var variable1 string = "Variable 1"
	fmt.Println(variable1)

	variable2 := "Varible 2"
	fmt.Println(variable2)

	var (
		variable3 string = "Varible 3"
		variable4 string = "Varible 4"
	)
	fmt.Println(variable3, variable4)

	variable5, variable6 := "Varible 5", "Varible 6"
	fmt.Println(variable5, variable6)

	const constant1 string = "Constant 1"
	fmt.Println(constant1)

	variable5, variable6 = variable6, variable5
	fmt.Println(variable5, variable6)
}
