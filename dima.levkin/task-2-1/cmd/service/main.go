package main

import (
	"fmt"
)

type Direction uint8

const (
	MoreOrEqual Direction = iota
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
		return 0, fmt.Errorf("bad operation: expected <= or >=, got %s", input)
	}

	return dir, nil
}

func checkBoundary(n uint8, low uint8, high uint8) bool {
	return n >= low && n <= high
}

func processQueries(workerCount int) error {
	var (
		maxBottom uint8 = tempBoundaryMin
		minTop    uint8 = tempBoundaryMax
	)

	for range workerCount {
		var operation string

		_, err := fmt.Scan(&operation)
		if err != nil {
			return fmt.Errorf("reading operation: %w", err)
		}

		direction, err := parseOp(operation)
		if err != nil {
			return err
		}

		var temp uint8

		_, err = fmt.Scan(&temp)
		if err != nil {
			return fmt.Errorf("invalid temperature input: %d", temp)
		}

		if !checkBoundary(temp, tempBoundaryMin, tempBoundaryMax) {
			return fmt.Errorf("temparature is valid but doesn't fit the boundaries: %d", temp)
		}

		switch direction {
		case MoreOrEqual:
			maxBottom = max(maxBottom, temp)
		case LessOrEqual:
			minTop = min(minTop, temp)
		}

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
		fmt.Println("invalid N parameter")

		return
	}

	for range departmentCount {
		_, err := fmt.Scan(&workerCount)
		if err != nil || workerCount > kLimitation {
			fmt.Println("invalid K parameter")

			return
		}

		err = processQueries(workerCount)
		if err != nil {
			fmt.Println(err.Error())

			return
		}
	}
}
