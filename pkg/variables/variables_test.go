package variables

import (
	"testing"

	"gotest.tools/assert"
)

func TestDeclareZeroValuedIntVar(t *testing.T) {
	var res int = DeclareZeroValuedIntVar()
	err := "Expected value is 0"

	assert.Equal(t, res, 0, err)
}

func TestDeclareTwoZeroValuedIntVar(t *testing.T) {
	var a, b int = DeclareTwoZeroValuedIntVars()
	err := "Expected value is 0"

	assert.Equal(t, a, 0, err)
	assert.Equal(t, b, 0, err)
}

func TestDeclareStrVarWithSpecifiedValue(t *testing.T) {
	var res string = DeclareStrVarWithSpecifiedValue()
	err := "Expected value is 'Some String'"

	assert.Equal(t, res, "Some String", err)
}

func TestDeclareVarsWithDifferentType(t *testing.T) {
	intVar1, intVar2, strVar, floatVar := DeclareVarsWithDifferentTypes()

	assert.Equal(t, intVar1, 0, "Expected value is 0")
	assert.Equal(t, intVar2, 0, "Expected value is 0")
	assert.Equal(t, strVar, "String", "Expected value is 'String'")
	assert.Equal(t, floatVar, 1.300000, "Expected value is 1.300000")
}

func TestDeclareVarInShortWay(t *testing.T) {
	var res string = DeclareVarInShortWay()
	err := "Expected value is `It was declared without provideing a type.`"

	assert.Equal(t, res, "It was declared without provideing a type.", err)

}
