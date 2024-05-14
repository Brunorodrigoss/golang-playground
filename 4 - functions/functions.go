package main

import "fmt"

func calcSum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {
	result := calcSum(10, 20)
	fmt.Println(result)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Function F")

	resultSum, resultSub := calcs(10, 15)
	fmt.Println(resultSum, resultSub)

	resSum, _ := calcs(30, 20)
	fmt.Println(resSum)

	_, resSub := calcs(30, 20)
	fmt.Println(resSub)
}
