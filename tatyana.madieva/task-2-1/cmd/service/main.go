package main

import (
	"fmt"
	"os"
)

func processDepartment(employees int) error {
	tMax := 30
	tMin := 15

	for range employees {
		var operation string
		var value int

		_, err := fmt.Scan(&operation, &value)
		if err != nil {
			return fmt.Errorf("reading operation: %w", err)
		}

		switch operation {
		case "<=":
			if tMax > value {
				tMax = value
			}
		case ">=":
			if tMin < value {
				tMin = value
			}
		default:
			return fmt.Errorf("unknown operation: %s", operation)
		}

		if tMin <= tMax {
			fmt.Println(tMin)
		} else {
			fmt.Println(-1)
		}
	}

	return nil
}

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading number of departments:", err)
		os.Exit(1)
	}

	if departments <= 0 {
		fmt.Fprintln(os.Stderr, "invalid number of departments:", departments)
		os.Exit(1)
	}

	for range departments {
		var employees int

		_, err = fmt.Scan(&employees)
		if err != nil {
			fmt.Fprintln(os.Stderr, "reading number of employees:", err)
			os.Exit(1)
		}

		if employees <= 0 {
			fmt.Fprintln(os.Stderr, "invalid number of employees:", employees)
			os.Exit(1)
		}

		err = processDepartment(employees)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
