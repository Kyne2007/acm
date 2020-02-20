package main

import "fmt"

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		fmt.Println(left)
		fmt.Println(right)

		mid := (left + right) / 2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
	}

	return numbers[left]
}
