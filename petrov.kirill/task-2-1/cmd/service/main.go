package main

import (
	"fmt"
)

func whattemp(left int, right int) {
	if left <= right {
		fmt.Println(left)
	} else {
		fmt.Println(-1)
	}
}

func main() {
	var num, kol int

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("errof of read num", err)

		return
	}

	for range num {
		_, err = fmt.Scan(&kol)
		if err != nil {
			fmt.Println("error of read kol", err)

			return
		}

		var left, right, add int

		var str string

		left = 15
		right = 30

		for range kol {
			_, err = fmt.Scan(&str, &add)
			if err != nil {
				fmt.Println("str && add reading error", err)

				return
			}

			if str[:2] == ">=" && add > left {
				left = add
			} else if str[:2] == "<=" && add < right {
				right = add
			}

			whattemp(left, right)
		}
	}
}
