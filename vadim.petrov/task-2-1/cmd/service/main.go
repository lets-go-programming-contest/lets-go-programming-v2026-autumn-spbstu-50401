package main

import "fmt"

func main() {
	var departatmentsCount int

	if _, err := fmt.Scanln(&departatmentsCount); err != nil {
		fmt.Println("Invalid read count of departaments(", err, ")")
		return
	}

	for i := 0; i < departatmentsCount; i++ {
		var peopleCount int

		if _, err := fmt.Scanln(&peopleCount); err != nil {
			fmt.Println("Invalid read count of people(", err, ")")
			return
		}

	}
}
