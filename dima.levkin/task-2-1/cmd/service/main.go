package main

import (
	"errors"
	"fmt"
)

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

var ErrInvalidOperator = errors.New("invalid operator")

func parseOp(input string) (Direction, error) {
	var dir Direction

	switch input {
	case ">=":
		dir = MoreOrEqual
	case "<=":
		dir = LessOrEqual
	default:
		return 0, ErrInvalidOperator
	}

	return dir, nil
}

func checkBoundary(n uint8, low uint8, high uint8) bool {
	return n >= low && n <= high
}

func proccessQueries(workerCount int) error {
	var (
		maxBottom uint8 = tempBoundaryMin
		minTop    uint8 = tempBoundaryMax
	)

	for range workerCount {
		var operation string

		_, err := fmt.Scan(&operation)
		if err != nil {
			fmt.Println("Invalid operation input")

			return ErrInvalidOperator
		}

		direction, err := parseOp(operation)
		if err != nil {
			fmt.Println("Invalid operation input parsing")

			return ErrInvalidOperator
		}

		var temp uint8

		_, err = fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Invalid temperature input")

			return ErrInvalidOperator
		}

		if !checkBoundary(temp, tempBoundaryMin, tempBoundaryMax) {
			fmt.Println("Invalid: temp not in boundary")

			return ErrInvalidOperator
		}

		switch direction {
		case MoreOrEqual:
			maxBottom = max(maxBottom, temp)
		case LessOrEqual:
			minTop = min(minTop, temp)
		} // either one of them will be called, I controlled for  that

		var result int

		if maxBottom > minTop {
			result = defaultResult
		} else {
			result = int(maxBottom)
		}

		fmt.Println(result)
	}

	return nil
}

func main() {
	var (
		departmentCount int
		workerCount     int
	)

	_, err := fmt.Scan(&departmentCount)
	if err != nil || departmentCount > nLimitation {
		fmt.Println("Invalin N parameter")

		return
	}

	for range departmentCount {
		_, err := fmt.Scan(&workerCount)
		if err != nil || workerCount > kLimitation {
			fmt.Println("Invalid K parameter")

			return
		}

		err = proccessQueries(workerCount)
		if err != nil {
			fmt.Println(err.Error())

			return
		}
	}
}
