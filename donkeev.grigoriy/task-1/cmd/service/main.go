package main

import "fmt"

func main() {
	var firstOp int
	_, err := fmt.Scanln(&firstOp)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var secondOp int
	_, err = fmt.Scanln(&secondOp)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operator string
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operator")
		return
	}

	switch operator {
	case "+":
		fmt.Println(firstOp + secondOp)
	case "-":
		fmt.Println(firstOp - secondOp)
	case "*":
		fmt.Println(firstOp * secondOp)
	case "/":
		if secondOp == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(firstOp / secondOp)
	default:
		fmt.Println("Invalid operation")
	}
}
