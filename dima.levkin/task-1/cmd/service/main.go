package main

import "fmt"

func main() {
	var (
		operation      string
		first_operand  int
		second_operand int
	)
	_, err := fmt.Scan(&first_operand)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&second_operand)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	if second_operand == 0 && operation == "/" {
		fmt.Println("Division by zero")
		return
	}

	var result = first_operand

	switch operation {
	case "+":
		result += second_operand
	case "-":
		result -= second_operand
	case "*":
		result *= second_operand
	case "/":
		result /= second_operand
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)

}
