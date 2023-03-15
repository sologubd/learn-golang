package main

import (
	"fmt"
	"pkg/variables"
)

func main() {
	fmt.Println("Let's get started!")

	PrintItem("Variables")
	variables.InitVariables()
}

func PrintItem(s string) {
	fmt.Println("####################")
	fmt.Printf("#     %s\n", s)
	fmt.Println("####################")
}
