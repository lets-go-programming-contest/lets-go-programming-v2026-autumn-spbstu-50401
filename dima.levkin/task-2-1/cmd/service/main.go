package main

import (
	"fmt"
)

type Temperature uint8
type Direction uint8

const (
	MoreOrEqual = iota
	LessOrEqual
)

func parseOp(input string) (Direction, error) {
	var dir Direction
	switch input {
	case ">=":
		dir = MoreOrEqual
	case "<=":
		dir = LessOrEqual
	default:
		return 0, fmt.Errorf("Invalid temperature input")

	}
	return dir, nil
}

func main() {
	const defaultResult = -1
	const (
		temperatureBoundaryMax = 30
		temperatureBoundaryMin = 15

		nLimitation = 1000
		kLimitation = 1000
	)

	var (
		n int
		k int
	)

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalin N parameter")
		return
	}

	for nI := 0; nI < n; nI++ {
		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Println("Invalin K parameter")
			return
		}

		directions := make([]Direction, 0, k)
		temperatures := make([]Temperature, 0, k)
		for kI := 0; kI < k; kI++ {
			var operation string

			_, err := fmt.Scan(&operation)
			if err != nil {
				fmt.Println("Invalid operation input")
				return
			}

			direction, err := parseOp(operation)
			if err != nil {
				fmt.Println("Invalid operation input parsing")
				return
			}

			var temperature Temperature
			_, err = fmt.Scan(&temperature)
			if err != nil {
				fmt.Println("Invalid temperature input")
				return
			}
			directions = append(directions, direction)
			temperatures = append(temperatures, temperature)
		}

	}
}
