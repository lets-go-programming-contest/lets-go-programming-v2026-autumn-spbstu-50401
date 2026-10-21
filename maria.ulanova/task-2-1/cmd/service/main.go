package main

import "fmt"

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println(err)

		return
	}

	for dep := departments; dep > 0; dep-- {
		var employees int

		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println(err)
			return
		}

		for emp := employees; emp > 0; emp-- {
			var (
				sign string
				temp int
			)

			if _, err := fmt.Scan(&sign, &temp); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(sign, temp)
		}
	}
}
