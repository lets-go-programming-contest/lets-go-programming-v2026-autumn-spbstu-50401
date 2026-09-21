package main

import (
	"fmt"
)

func main() {
	var num1 int
	_, err1 := fmt.Scan(&num1)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var num2 int
	_, err2 := fmt.Scan(&num2)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var sign string
	_, errSign := fmt.Scan(&sign)
	if errSign != nil || len(sign) != 1 {
		fmt.Println("Invalid operation")
		return
	}

	var result int
	result = num1

	switch sign {
	case "+":
		result += num2
	case "-":
		result -= num2
	case "*":
		result *= num2
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero")
		}
		result /= num2
	default:
		fmt.Println("Invalid operation")
		return

	}
	fmt.Println(result)

}
