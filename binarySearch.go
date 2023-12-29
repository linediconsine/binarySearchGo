package main

import (
	"fmt"
)

func binarySort(list []int64, target int64) (position int64) {
	max := len(list) - 1
	min := 0

	for min <= max {
		mid := min + (max-min)/2

		if list[mid] == target {
			fmt.Println("OK")
			return int64(mid)
		}

		if target < list[mid] {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}

func main() {
	array := []int64{1, 3, 4, 5}

	fmt.Println(binarySort(array, 5))
}
