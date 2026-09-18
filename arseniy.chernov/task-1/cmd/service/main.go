package main

import (
	"fmt"
)

func main() {
	var a, b int
	var op string
	n, err := fmt.Scan(&a, &b, &op)
	if err != nil && n == 0 {
		fmt.Println("Invalid first operand")
		return
	} else if err != nil && n == 1 {
		fmt.Println("Invalid second operand")
		return
	}

	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = a / b
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}
