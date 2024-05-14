package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
	city   string
}

func main() {
	var user1 user
	user1.name = "Bruno"
	user1.age = 32
	fmt.Println(user1)

	address := address{"Xpto Street", 0, "City"}

	user2 := user{"Rodrigo", 30, address}
	fmt.Println(user2)

	user3 := user{name: "Luis"}
	fmt.Println(user3)
}
