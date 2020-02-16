package main

import "testing"

func TestAddition(t *testing.T) {
	if calc("2+2") != 4 {
		t.Errorf("test addition failed")
	}
	if calc("(1+2)+(3+4)") != 10 {
		t.Errorf("test addition with brackets failed")
	}
}

func TestSubtraction(t *testing.T) {
	if calc("10-4") != 6 {
		t.Errorf("test subtraction failed")
	}
	if calc("(4-1)-(3-2)") != 2 {
		t.Errorf("test subtraction with brackets failed")
	}
}

func TestMultiplication(t *testing.T) {
	if calc("4*5") != 20 {
		t.Errorf("test multiplication failed")
	}
	if calc("(1*2)*(3*4)") != 24 {
		t.Errorf("test multiplication with brackets failed")
	}
}

func TestDivision(t *testing.T) {
	if calc("66/6") != 11 {
		t.Errorf("test division failed")
	}
	if calc("(6/2)/(2/2)") != 3 {
		t.Errorf("test division with brackets failed")
	}
}

func TestCalc(t *testing.T) {
	if calc("(1+2)-3") != 0 {
		t.Errorf("test calc failed")
	}
	if calc("(1+2)*3") != 9 {
		t.Errorf("test calc failed")
	}
}
