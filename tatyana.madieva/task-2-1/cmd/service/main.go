package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)
		var tMax = 30
		var tMin = 15
		for j := 0; j < k; j++ {
			var op string
			var v int
			fmt.Scan(&op, &v)
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
				fmt.Println("error")
			}

			if tMin <= tMax {
				fmt.Println(tMin)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
