package main

import (
	"fmt"
)

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)

	array := make([]int, n+1)
	for i := 0; i < m; i++ {
		l, r := 0, 0
		fmt.Scan(&l, &r)

		array[l-1]++
		array[r]--
	}

	for i := 1; i < n; i++ {
		array[i] += array[i-1]
	}

	min := 200000
	for i := 0; i < n; i++ {
		if array[i] < min {
			min = array[i]
			//fmt.Printf("array[%d] = %d\n", i, array[i])
		}
	}
	fmt.Println(min)
}
