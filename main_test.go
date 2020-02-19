package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func M(a, b interface{}) []interface{} {
	return []interface{}{a, b}
}

func TestAddition(t *testing.T) {
	require.Equal(t, M(calc("2+2")), M(4.0, nil),
		"test addition failed")
	require.Equal(t, M(calc("(1+2)+(3+4)")), M(10.0, nil),
		"test addition with brackets failed")
}

func TestSubtraction(t *testing.T) {
	require.Equal(t, M(calc("10-4")), M(6.0, nil),
		"test subtraction failed")
	require.Equal(t, M(calc("(4-1)-(3-2)")), M(2.0, nil),
		"test subtraction with brackets failed")
}

func TestMultiplication(t *testing.T) {
	require.Equal(t, M(calc("4*5")), M(20.0, nil),
		"test multiplication failed")
	require.Equal(t, M(calc("(1*2)*(3*4)")), M(24.0, nil),
		"test multiplication with brackets failed")
}

func TestDivision(t *testing.T) {
	require.Equal(t, M(calc("66/6")), M(11.0, nil),
		"test division failed")
	require.Equal(t, M(calc("(6/2)/(2/2)")), M(3.0, nil),
		"test division with brackets failed")
}

func TestCalc(t *testing.T) {
	require.Equal(t, M(calc("(1+2)-3")), M(0.0, nil),
		"test calc failed")
	require.Equal(t, M(calc("(1+2)*3")), M(9.0, nil),
		"test calc failed")
}

func TestAddingLessPriorityOperator(t *testing.T) {
	require.Equal(t, M(calc("2*2-1")), M(3.0, nil),
		"test adding less priority operator failed")
	require.Equal(t, M(calc("(2+2)*10-7")), M(33.0, nil),
		"test adding less priority operator with brackets failed")
}

func TestCalcWithLongNumber(t *testing.T) {
	require.Equal(t, M(calc("11+22*3")), M(77.0, nil),
		"test calc with long number failed")
	require.Equal(t, M(calc("(11+22)*3")), M(99.0, nil),
		"test calc with long number and brackets failed")
}

func TestCalcWithInvalidBrackets(t *testing.T) {
	_, err := calc("1+1)")
	require.Equal(t, err.Error(), "invalid brackets", "test calc with invalid brackets failed")
}

func TestCalcWithInvalidSymbol(t *testing.T) {
	_, err := calc("1+1#2")
	require.Equal(t, err.Error(), "invalid symbol", "test calc with invalid symbol failed")
}
