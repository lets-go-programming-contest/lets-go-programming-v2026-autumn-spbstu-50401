package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	startMinTemp = 15
	startMaxTemp = 30
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var departmentCount int

	if _, err := fmt.Fscan(reader, &departmentCount); err != nil {
		fmt.Println(err)

		return
	}

	for range departmentCount {
		var employeeCount int

		if _, err := fmt.Fscan(reader, &employeeCount); err != nil {
			fmt.Println(err)

			return
		}

		minTemp := startMinTemp
		maxTemp := startMaxTemp

		for range employeeCount {
			var (
				operation   string
				temperature int
			)

			if _, err := fmt.Fscan(reader, &operation, &temperature); err != nil {
				fmt.Println(err)

				return
			}

			switch operation {
			case ">=":
				if temperature > minTemp {
					minTemp = temperature
				}
			case "<=":
				if temperature < maxTemp {
					maxTemp = temperature
				}
			default:
				fmt.Println("invalid sign:", operation)

				return
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
