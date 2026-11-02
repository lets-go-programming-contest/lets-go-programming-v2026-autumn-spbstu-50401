package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, val)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]

	return item
}

func main() {
	var (
		numberOfDishes int
		dishRating     int
		ratingChoice   int
	)

	_, err := fmt.Scan(&numberOfDishes)
	if err != nil {
		fmt.Printf("failed to scan number of dishes: %v\n", err)

		return
	}

	intHeap := make(IntHeap, 0, numberOfDishes)
	heap.Init(&intHeap)

	for range numberOfDishes {
		_, err := fmt.Scan(&dishRating)
		if err != nil {
			fmt.Printf("failed to scan dish rating: %v\n", err)

			return
		}

		heap.Push(&intHeap, dishRating)
	}

	_, err = fmt.Scan(&ratingChoice)
	if err != nil {
		fmt.Printf("failed to scan rating choice: %v\n", err)

		return
	}

	if ratingChoice <= 0 || ratingChoice > numberOfDishes {
		fmt.Printf("invalid rating choice: must be between 1 and %d\n", numberOfDishes)

		return
	}

	for range ratingChoice - 1 {
		heap.Pop(&intHeap)
	}

	fmt.Println(heap.Pop(&intHeap))
}
