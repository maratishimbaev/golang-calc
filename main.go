package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getPriority(symbol string) int {
	switch true {
	case strings.Index("*/", symbol) != -1:
		return 3
	case strings.Index("+-", symbol) != -1:
		return 2
	case strings.Index("(", symbol) != -1:
		return 1
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
		case strings.Index("1234567890", symbol) != -1:
			if rpnStr != "" && strings.Index("1234567890", prevSymbol) == -1 {
				rpnStr += " "
			}
			rpnStr += symbol
		case strings.Index("+-*/", symbol) != -1:
			if len(stack) == 0 || getPriority(stack[len(stack) - 1]) < getPriority(symbol) {
				stack = append(stack, symbol)
			} else {
				for len(stack) > 0 && getPriority(stack[len(stack) - 1]) >= getPriority(symbol) {
					rpnStr = rpnStr + " " + stack[len(stack) - 1]
					stack = stack[:len(stack) - 1]
				}
			}
		case strings.Index("(", symbol) != -1:
			stack = append(stack, symbol)
		case strings.Index(")", symbol) != -1:
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

func calc(str string) int {
	rpnStr := transformToRpn(str)

	stack := make([]int, 0, 10)
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
			stack = append(stack, num)
		}
	}

	return stack[len(stack) - 1]
}

func main() {
	fmt.Println(calc("(1-2)*3"))
}