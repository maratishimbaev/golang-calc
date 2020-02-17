package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	require.Equal(t, calc("2+2"), 4.0,
		"test addition failed")
	require.Equal(t, calc("(1+2)+(3+4)"), 10.0,
		"test addition with brackets failed")
}

func TestSubtraction(t *testing.T) {
	require.Equal(t, calc("10-4"), 6.0,
		"test subtraction failed")
	require.Equal(t, calc("(4-1)-(3-2)"), 2.0,
		"test subtraction with brackets failed")
}

func TestMultiplication(t *testing.T) {
	require.Equal(t, calc("4*5"), 20.0,
		"test multiplication failed")
	require.Equal(t, calc("(1*2)*(3*4)"), 24.0,
		"test multiplication with brackets failed")
}

func TestDivision(t *testing.T) {
	require.Equal(t, calc("66/6"), 11.0,
		"test division failed")
	require.Equal(t, calc("(6/2)/(2/2)"), 3.0,
		"test division with brackets failed")
}

func TestCalc(t *testing.T) {
	require.Equal(t, calc("(1+2)-3"), 0.0,
		"test calc failed")
	require.Equal(t, calc("(1+2)*3"), 9.0,
		"test calc failed")
}

func TestAddingLessPriorityOperator(t *testing.T) {
	require.Equal(t, calc("2*2-1"), 3.0,
		"test adding less priority operator failed")
	require.Equal(t, calc("(2+2)*10-7"), 33.0,
		"test adding less priority operator with brackets failed")
}

func TestCalcWithLongNumber(t *testing.T) {
	require.Equal(t, calc("11+22*3"), 77.0,
		"test calc with long number failed")
	require.Equal(t, calc("(11+22)*3"), 99.0,
		"test calc with long number and brackets failed")
}
