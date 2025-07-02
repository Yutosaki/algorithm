package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	max := 0
	for i := 0; i < n; i++ {
		if a[i] >= i+1 {
			max = i + 1
		} else {
			break
		}
	}
	fmt.Println(max)
}
