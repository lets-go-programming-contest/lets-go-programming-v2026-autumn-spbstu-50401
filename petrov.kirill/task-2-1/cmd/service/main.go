package main

import (
	"fmt"
)

func main() {
	var num, kol int

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}

	for range num {
		_, err = fmt.Scan(&kol)
		if err != nil {
			fmt.Println(err)
		}

		var left, right, add int

		var str string

		left = 15
		right = 30

		for range kol {
			_, err = fmt.Scan(&str, &add)
			if err != nil {
				fmt.Println(err)
			}

			if string(str[0]) == ">" && add > left {
				left = add
			} else if string(str[0]) == "<" && add < right {
				right = add
			}

			if left <= right {
				fmt.Println(left)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
