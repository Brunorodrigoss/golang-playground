package main

import (
	"fmt"
	"learnPackages/auxiliary"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Print from main")
	auxiliary.Write()

	err := checkmail.ValidateFormat("bruno@gmail.com")
	fmt.Println(err)

}
