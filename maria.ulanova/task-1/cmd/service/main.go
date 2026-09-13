package main

import "fmt"

func main() {
	var a, b int
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	fmt.Println(a, b)
}
