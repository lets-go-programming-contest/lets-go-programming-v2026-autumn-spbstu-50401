package main

import (
	"errors"
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

var (
	ErrBadOperation           = errors.New("bad operation: expected <= or >=")
	ErrInvalidTemperature     = errors.New("invalid temperature input")
	ErrTemperatureOutOfBounds = errors.New("temperature is valid but doesn't fit the boundaries")
)

func parseOp(input string) (Direction, error) {
	var dir Direction

	switch input {
	case ">=":
		dir = MoreOrEqual
	case "<=":
		dir = LessOrEqual
	default:
		return 0, fmt.Errorf("%w, got %s", ErrBadOperation, input)
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

		var temperature uint8

		_, err = fmt.Scan(&temperature)
		if err != nil {
			return fmt.Errorf("%w: %d", ErrInvalidTemperature, temperature)
		}

		if !checkBoundary(temperature, tempBoundaryMin, tempBoundaryMax) {
			return fmt.Errorf("%w: %d", ErrTemperatureOutOfBounds, temperature)
		}

		switch direction {
		case MoreOrEqual:
			maxBottom = max(maxBottom, temperature)
		case LessOrEqual:
			minTop = min(minTop, temperature)
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
