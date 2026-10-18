package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Invalid department count")
		return
	}

	var k int
	if _, err := fmt.Scan(&k); err != nil {
		fmt.Println("Invalid employee count")
		return
	}

	for i := 0; i < n; i++ {
		var minTemp int = 15
		var maxTemp int = 30

		for j := 0; j < k; i++ {
			var op string
			var t int

			if _, err := fmt.Scan(&op, &t); err != nil {
				fmt.Println("Invalid input temperature data")
				return
			}

		}

	}
}
