package variables

/*
Displays different ways to declare variables

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

/*
To declare a variable we can use the keyword `var`
also the type should be provided, each type has a default value
*/
func DeclareZeroValuedIntVar() int {
	var intVar int
	fmt.Printf("`intVar`, default value: %d\n", intVar)

	return intVar
}

/*
A few variables might be declared at the same time
*/
func DeclareTwoZeroValuedIntVars() (a, b int) {
	var intVar1, intVar2 int

	fmt.Printf("`intVar1`, default value: %d\n", intVar1)
	fmt.Printf("`intVar2`, default value: %d\n", intVar2)

	a, b = intVar1, intVar2

	return
}

/*
A value might be initialized.
If the value is provided, the type might be omitted
*/
func DeclareStrVarWithSpecifiedValue() string {
	var strVar = "Some String"
	fmt.Printf("`strVar`, value: %s\n", strVar)
	return strVar
}

/*
Multilined `var` declaration is also available
*/
func DeclareVarsWithDifferentTypes() (a, b int, c string, d float64) {
	var (
		intVar1, intVar2 int
		strVar           = "String"
		floatVar         = 1.3
	)

	a, b, c, d = intVar1, intVar2, strVar, floatVar

	fmt.Printf("`a`, value: %d\n", a)
	fmt.Printf("`b`, value: %d\n", b)
	fmt.Printf("`c`, value: %s\n", c)
	fmt.Printf("`d`, value: %f\n", d)

	return
}

/*
Short variable declaration is also available
It requires neither providing a type nor using a keyword `var`
*/
func DeclareVarInShortWay() string {
	shortStrVar := "It was declared without provideing a type."

	fmt.Printf("`shortStrVar, value: %s\n", shortStrVar)

	return shortStrVar
}
