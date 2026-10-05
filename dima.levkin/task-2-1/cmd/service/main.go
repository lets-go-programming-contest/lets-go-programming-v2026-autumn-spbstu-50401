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
const defaultResult = -1
const (
	tempBoundaryMax = 30
	tempBoundaryMin = 15

	nLimitation = 1000
	kLimitation = 1000
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
func checkBoundary(n uint8, low uint8, high uint8) bool {
	return n >= low && n <= high
}

func main() {

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


		var (
			maxBottom uint8 = 15
			minTop    uint8 = 30
		)
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

			var temp uint8
			_, err = fmt.Scan(&temp)
			if err != nil {
				fmt.Println("Invalid temperature input")
				return
			}
			if false == checkBoundary(temp, tempBoundaryMin, tempBoundaryMax) {
				fmt.Println("Invalid: temp not in boundary")
				return
			}

			switch direction {
			case MoreOrEqual:
				maxBottom = max(maxBottom, temp)
			case LessOrEqual:
				minTop = min(minTop, temp)
			} // either one of them will be called, I controlled for  that

			var result int = 0
			if maxBottom > minTop {
				result = defaultResult
			} else {
				result = int(maxBottom)
			}
			fmt.Println(result)
		}
	
	}

}
