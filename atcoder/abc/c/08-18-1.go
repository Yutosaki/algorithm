package main

import (
	"fmt"
)

type nums struct {
	characterLen int
	x            int
	beforeLen    int
}

func binarySearch(array []nums, key int) int {
	left := -1
	right := len(array)

	for right-left > 1 {
		mid := left + (right-left)/2

		if array[mid].beforeLen <= key {
			left = mid
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	q := 0
	fmt.Scan(&q)
	totalLen := 0
	array := make([]nums, 0)
	for q > 0 {
		q--
		taken := 0
		fmt.Scan(&taken)

		if taken == 1 {
			var newNum nums
			fmt.Scan(&newNum.characterLen, &newNum.x)

			newNum.beforeLen = totalLen
			totalLen += newNum.characterLen

			array = append(array, newNum)
		} else {
			k := 0
			fmt.Scan(&k)

			sum := 0
			k--
			index := binarySearch(array, k)
			for i := 0; i < index; i++ {
				sum += array[i].x * array[i].characterLen
			}

			rest := k - array[index].beforeLen + 1
			array[index].characterLen -= rest

			sum += rest * array[index].x
			fmt.Println(sum)

			array = array[index:]

			totalLen = 0
			for i := 0; i < len(array); i++ {
				array[i].beforeLen = totalLen
				totalLen += array[i].characterLen
			}
		}
	}
}
