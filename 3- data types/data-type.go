package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64
	var number int64 = -100000000
	fmt.Println(number)

	//unsigned integer
	var number2 uint32 = 10000
	fmt.Println(number2)

	// Alias
	// uint8 = byte
	// int32 = rune

	var number3 rune = 123456
	fmt.Println(number3)

	var number4 byte = 123
	fmt.Println(number4)

	//float32, float64
	var realNumber1 float32 = 123.99
	fmt.Println(realNumber1)

	var realNumber2 float64 = 1230000000000000000.99
	fmt.Println(realNumber2)

	var str string = "Text"
	fmt.Println(str)

	str2 := "Text 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var text string
	fmt.Println(text)

	var value int
	fmt.Println(value)

	var boolean1 bool = true
	fmt.Println(boolean1)

	var boolean2 bool
	fmt.Println(boolean2)

	var err error
	fmt.Println(err)

	var err2 error = errors.New("Internal error")
	fmt.Println(err2)

}
