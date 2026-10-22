package main

import "fmt"

func main() {
	var lhs, rhs int
	var op string

	_, err1 := fmt.Scan(&lhs)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err2 := fmt.Scan(&rhs)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err3 := fmt.Scan(&op)
	if err3 != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch op {
	case "+":
		fmt.Println(lhs + rhs)
	case "-":
		fmt.Println(lhs - rhs)
	case "*":
		fmt.Println(lhs * rhs)
	case "/":
		if rhs == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(lhs / rhs)
	default:
		fmt.Println("Invalid operation")
	}
}
