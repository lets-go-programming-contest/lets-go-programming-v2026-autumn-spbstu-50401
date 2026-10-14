package main

import "fmt"

func main() {
	var N, K int

	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Ошибка чтения")
		return
	}
}
