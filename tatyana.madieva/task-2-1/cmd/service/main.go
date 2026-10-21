package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading number of departments:", err)
		os.Exit(1)
	}
	if n <= 0 {
		fmt.Fprintln(os.Stderr, "invalid number of departments:", n)
		os.Exit(1)
	}
	for i := 0; i < n; i++ {
		var k int
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Fprintln(os.Stderr, "reading number of employees:", err)
			os.Exit(1)
		}
		if k <= 0 {
			fmt.Fprintln(os.Stderr, "invalid number of employees:", k)
			os.Exit(1)
		}
		var tMax = 30
		var tMin = 15
		for j := 0; j < k; j++ {
			var op string
			var v int
			_, err = fmt.Scan(&op, &v)
			if err != nil {
				fmt.Fprintln(os.Stderr, "reading operation:", err)
				os.Exit(1)
			}
			switch op {
			case "<=":
				if tMax > v {
					tMax = v
				}
			case ">=":
				if tMin < v {
					tMin = v
				}
			default:
				fmt.Fprintln(os.Stderr, "unknown operation:", op)
				os.Exit(1)
			}

			if tMin <= tMax {
				fmt.Println(tMin)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
