package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	d := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		fmt.Scan(&d[i])
	}

	total := 0
	for i := 0; i < n-1; i++ {
		total = 0
		for j := i; j < n-1; j++ {
			fmt.Print(d[j]+total, " ")
			total += d[j]
		}
		fmt.Println()
	}
}
