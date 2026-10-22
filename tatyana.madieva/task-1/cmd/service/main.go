package main

import "fmt"

func main() {
	var value1 int
	if _, err := fmt.Scan(&value1); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var value2 int
	if _, err := fmt.Scan(&value2); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var sOp string
	if _, err := fmt.Scan(&sOp); err != nil {
		fmt.Println("Invalid operation")
		return
	}
	var res int
	switch sOp {
	case "+":
		res = value1 + value2
	case "-":
		res = value1 - value2
	case "*":
		res = value1 * value2
	case "/":
		if value2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		res = value1 / value2
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(res)
}
