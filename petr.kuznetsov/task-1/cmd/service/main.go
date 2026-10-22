package main

import "fmt"

func main() {
	var a, b int
	var op rune

	_, efirst := fmt.Scan(&a)
	if efirst != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, esecond := fmt.Scan(&b)
	if esecond != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err := fmt.Scanf("%c\n", &op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch op {
	case '+':
		fmt.Println(a + b)
	case '-':
		fmt.Println(a - b)
	case '*':
		fmt.Println(a * b)
	case '/':
		if b == 0 {
			fmt.Println("Division by zero")
			return
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
