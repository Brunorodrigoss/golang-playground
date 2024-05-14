package main

import "fmt"

func main() {
	var variable1 string = "variable1"
	variable2 := "variable2"
	fmt.Println(variable1)
	fmt.Println(variable2)

	// Arithmetics
	sum := 1 + 2
	subtraction := 1 - 2
	division := 10 / 4
	multiplication := 10 * 5
	module := 10 % 2

	fmt.Println(sum, subtraction, division, multiplication, module)

	var number1 int16 = 10
	var number2 int32 = 25

	result := number1 + int16(number2)
	fmt.Println(result)

	// Relational
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)

	// Logics

	fmt.Println(true && false) // and
	fmt.Println(true || false) // or
	fmt.Println(!true)         // not
	fmt.Println(!false)        // not

	// Unary
	number := 10
	number++
	fmt.Println(number)

	number += 15
	fmt.Println(number)

	number--
	fmt.Println(number)

	number -= 15
	fmt.Println(number)

	number *= 10
	fmt.Println(number)

	number /= 2
	fmt.Println(number)

	number %= 5
	fmt.Println(number)
}
