package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type student struct {
	person
	course  string
	college string
}

func main() {
	person1 := person{"Bruno", 32}
	student := student{person1, "Computer Science", "UNICAMP"}

	fmt.Println(student)
}
