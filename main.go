package main

import (
	"errors"
	"fmt"
)

var NotFound error = errors.New("Не найдено")

func binarySearchs(s []int, item int) (int, error) {
	low := 0
	high := len(s) - 1
	mid := -1
	for low <= high {
		mid = (high + low) / 2
		if s[mid] == item {
			return mid, nil
		}
		if s[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, NotFound
}

func main() {
	s := []int{1, 4, 6, 7, 8}
	answer, err := binarySearchs(s, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer)
}
