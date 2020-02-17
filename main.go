package main

import (
	"fmt"
	"strconv"
	"strings"
)

var numberSet = "1234567890"
var operatorSet = "+-*/"

func getPriority(symbol string) int {
	switch symbol {
	case "*": return 3
	case "/": return 3
	case "+": return 2
	case "-": return 2
	case "(": return 1
	}
	return 0
}

func transformToRpn(str string) string {
	var rpnStr string
	stack := make([]string, 0, 10)

	var prevSymbol string

	for _, char := range str {
		symbol := string(char)

		switch true {
		case strings.Index(numberSet, symbol) != -1:
			if rpnStr != "" && strings.Index(numberSet, prevSymbol) == -1 {
				rpnStr += " "
			}
			rpnStr += symbol
		case strings.Index(operatorSet, symbol) != -1:
			if len(stack) == 0 || getPriority(stack[len(stack) - 1]) < getPriority(symbol) {
				stack = append(stack, symbol)
			} else {
				for len(stack) > 0 && getPriority(stack[len(stack) - 1]) >= getPriority(symbol) {
					rpnStr = rpnStr + " " + stack[len(stack) - 1]
					stack = stack[:len(stack) - 1]
				}
				stack = append(stack, symbol)
			}
		case symbol == "(":
			stack = append(stack, symbol)
		case symbol == ")":
			for len(stack) > 0 && stack[len(stack) - 1] != "(" {
				rpnStr = rpnStr + " " + stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
			}
			stack = stack[:len(stack) - 1]
		}

		prevSymbol = symbol
	}

	for len(stack) > 0 {
		rpnStr = rpnStr + " " + stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
	}

	return rpnStr
}

func calc(str string) float64 {
	rpnStr := transformToRpn(str)

	stack := make([]float64, 0, 10)
	expArray := strings.Split(rpnStr, " ")

	for _, val := range expArray {
		if strings.Index("+-*/", val) != -1 {
			preLast := stack[len(stack) - 2]
			last := stack[len(stack) - 1]

			switch val {
			case "+":
				stack[len(stack) - 2] = preLast + last
			case "-":
				stack[len(stack) - 2] = preLast - last
			case "*":
				stack[len(stack) - 2] = preLast * last
			case "/":
				stack[len(stack) - 2] = preLast / last
			}

			stack = stack[:len(stack) - 1]
		} else {
			num, _ := strconv.Atoi(val)
			floatNum := float64(num)
			stack = append(stack, floatNum)
		}
	}

	return stack[len(stack) - 1]
}

func main() {
	fmt.Println(calc("2*2-1"))
}