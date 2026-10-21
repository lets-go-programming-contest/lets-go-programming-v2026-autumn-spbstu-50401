package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

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

		lower, upper := minTemp, maxTemp

		for emp := employees; emp > 0; emp-- {
			var (
				sign string
				temp int
			)

			if _, err := fmt.Scan(&sign, &temp); err != nil {
				fmt.Println(err)

				return
			}

			switch sign {
			case ">=":
				if temp > lower {
					lower = temp
				}
			case "<=":
				if temp < upper {
					upper = temp
				}
			default:
				fmt.Println("invalid sign:", sign)

				return
			}

			if lower <= upper {
				fmt.Println(lower)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
