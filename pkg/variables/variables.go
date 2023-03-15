package variables

/*
Specification:
    - Variables
        - https://go.dev/ref/spec#Variables
    - Variables declaration
        - https://go.dev/ref/spec#Variable_declarations
    - Short variables declaration
        - https://go.dev/ref/spec#Short_variable_declarations
Tutorial:
    - https://go.dev/tour/basics/8
Go by examples
    - https://gobyexample.com/variables
*/

import (
	"fmt"
)

// Displays different ways to declare variables
func InitVariables() {
	/*
	   To declare a variable we can use the keyword `var`
	   also the type should be provided, each type has a default value
	*/
	var intVar int
	fmt.Printf("`intVar`, default value: %d\n", intVar)

	/*
	   A few variables might be declared at the same time
	*/
	var intVar1, intVar2 int
	fmt.Printf("`intVar1`, default value: %d\n", intVar1)
	fmt.Printf("`intVar2`, default value: %d\n", intVar2)

	/*
	   A value might be initialized.
	   If the value is provided, the type might be omitted
	*/
	var strVar = "Some String"
	fmt.Printf("`strVar`, value: %s\n", strVar)

	/*
	   Multilined `var` declaration is also available
	*/
	var (
		a, b int
		c    = "String"
		d    = 1.3
	)

	fmt.Printf("`a`, value: %d\n", a)
	fmt.Printf("`b`, value: %d\n", b)
	fmt.Printf("`c`, value: %s\n", c)
	fmt.Printf("`d`, value: %f\n", d)

	/*
	   Short variable declaration is also available
	   It requires neither providing a type nor using a keyword `var`
	*/
	shortStrVar := "It was declared without provideing a type."
	fmt.Printf("`shortStrVar, value: %s\n", shortStrVar)
}
